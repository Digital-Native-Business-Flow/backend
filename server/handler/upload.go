package handler

import (
	"encoding/csv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"

	"backend/ent"
	"backend/ent/instrument"
	"backend/internal"
)

type DateTime struct {
	time.Time
}

// UnmarshalCSV converts the CSV string to a usable date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("1/2/2006 15:04", csv)
	return err
}

type CSVData struct {
	Code uint     `csv:"ErrorCode"`
	Text string   `csv:"ErrorText"`
	Time DateTime `csv:"TestDateTime"`
}

func (h *Handler) Upload(c *fiber.Ctx) (err error) {
	// Parse the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return internal.NewError(internal.ErrBEMissingForm, err, 1)
	}

	// Get all files from the "documents" key
	files := form.File["documents"]

	// Loop through the files
	for _, file := range files {
		// Check the file type
		if file.Header["Content-Type"][0] != "text/csv" {
			return internal.NewError(internal.ErrBENotCSV, err, 1)
		}

		// Open the file for reading
		openedFile, err := file.Open()
		if err != nil {
			return internal.NewError(internal.ErrBEFileOpen, err, 1)
		}

		// Parse the instrument name
		instrumentName := strings.Split(file.Filename, "-")[0]

		// Create &/or select the ID of the instrument from DB
		i, err := h.createSelectInstrument(c, instrumentName)
		if err != nil {
			return err // Already processed
		}

		// Open CSV file reader
		csvFile := csv.NewReader(openedFile)

		// Set the comma (separator) type of the file
		csvFile.Comma = ';'

		// Unmarshal the CSV file
		var data []*CSVData
		if err := gocsv.UnmarshalCSV(csvFile, &data); err != nil {
			return internal.NewError(internal.ErrBECSVUnmarshal, err, 1)
		}

		// Close the file
		err = openedFile.Close()
		if err != nil {
			return internal.NewError(internal.ErrBEFileClose, err, 1)
		}

		// Loop through all entries and insert relevant data in DB
		for _, entry := range data {
			if entry.Code == 122 || entry.Code == 313 {
				_, err = h.DB.InstrumentError.
					Create().
					SetInstrumentID(i.ID).
					SetCode(entry.Code).
					SetText(entry.Text).
					SetOccurredAt(entry.Time.Time).
					Save(c.Context())
				if err != nil {
					return internal.NewError(internal.ErrDBInsert, err, 1)
				}
			}
		}
	}

	return nil
}

func (h *Handler) createSelectInstrument(c *fiber.Ctx, name string) (i *ent.Instrument, err error) {
	// Set the creation state
	create := false

	// Try and get the instrument ID using the given name
	i, err = h.DB.Instrument.
		Query().
		Select(
			instrument.FieldID,
		).
		Where(instrument.NameEQ(name)).
		Only(c.Context())
	if err != nil {
		// Build the error response based on error type
		switch err.(type) {
		case *ent.NotFoundError:
			// Nothing to do, means we need to create it
			create = true
		default:
			return nil, internal.NewError(internal.ErrDBQuery, err, 1)
		}
	}

	// Return the instrument data if found already
	if !create {
		return i, nil
	}

	// Create the new instrument
	i, err = h.DB.Instrument.
		Create().
		SetName(name).
		Save(c.Context())
	if err != nil {
		return nil, internal.NewError(internal.ErrDBInsert, err, 1)
	}

	return i, nil
}
