package chacha20poly1305

import (
	"reflect"
	"strings"
	"testing"
)

func TestPrepareInput(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		wantHeader  string
		wantData    []byte
		wantErr     bool
		expectedErr string // Add this field
	}{
		{
			name:       "Valid Input",
			input:      []byte("$HEADER\n" + "0102030405060708\n" + "090a0b0c0d0e0f10"),
			wantHeader: "$HEADER",
			wantData:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			wantErr:    false,
		},
		{
			name:        "Invalid Format",
			input:       []byte("HEADER_ONLY"),
			wantHeader:  "",
			wantData:    nil,
			wantErr:     true,
			expectedErr: "invalid input", //Check the error message
		},
		{
			name:        "Invalid Hex Data",
			input:       []byte("$HEADER\n" + "INVALID_HEX"),
			wantHeader:  "",
			wantData:    nil,
			wantErr:     true,
			expectedErr: "invalid byte", //Check the error message
		},
		{
			name:       "Empty Data",
			input:      []byte("$HEADER\n"),
			wantHeader: "$HEADER",
			wantData:   []byte{},
			wantErr:    false,
		},
		// {
		// 	name:       "With Whitespace",
		// 	input:      []byte("  $HEADER  \n  0102  \n  0304  \n"),
		// 	wantHeader: "$HEADER",
		// 	wantData:   []byte{0x01, 0x02, 0x03, 0x04},
		// 	wantErr:    false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, data, err := prepareInput(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("prepareInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !strings.Contains(err.Error(), tt.expectedErr) {
					t.Errorf("prepareInput() error message = %q, must contain %q", err.Error(), tt.expectedErr)
				}
				return // Stop here if we expect an error
			}

			if header != tt.wantHeader {
				t.Errorf("prepareInput() header = %v, want %v", header, tt.wantHeader)
			}
			if !reflect.DeepEqual(data, tt.wantData) {
				t.Errorf("prepareInput() data = %v, want %v", data, tt.wantData)
			}
		})
	}
}
