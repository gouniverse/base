package shared

import (
	"reflect"
	"strings"
	"testing"
)

func TestTextContainerParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantHeader  string
		wantBody    string
		wantErr     bool
		expectedErr string // Add this field
	}{
		{
			name:       "01. Valid Input",
			input:      "$HEADER\n" + "0102030405060708\n" + "090a0b0c0d0e0f10",
			wantHeader: "$HEADER",
			wantBody:   "0102030405060708090a0b0c0d0e0f10",
			wantErr:    false,
		},
		{
			name:        "02. Valid Format (Header Only)",
			input:       "HEADER_ONLY",
			wantHeader:  "HEADER_ONLY",
			wantBody:    "",
			wantErr:     false,
			expectedErr: "invalid characters in header", //Check the error message
		},
		{
			name:       "03. Valid Format (Header Only with Newline)",
			input:      "$HEADER\n",
			wantHeader: "$HEADER",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "04. Valid Format (Header and Whitespace)",
			input:      "$HEADER\n" + "\n\n \n \n",
			wantHeader: "$HEADER",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "05. Valid Format (Whitespace, Header and Body)",
			input:      "  $HEADER  \n  0102  \n\t  0304  \n",
			wantHeader: "$HEADER",
			wantBody:   "01020304",
			wantErr:    false,
		},
		{
			name:       "06. Valid Input",
			input:      "$HEADER\n" + "0102030405060708\n" + "090a0b0c0d0e0f10",
			wantHeader: "$HEADER",
			wantBody:   "0102030405060708090a0b0c0d0e0f10",
			wantErr:    false,
		},
		{
			name:       "07. Valid Format (Header Only)",
			input:      "HEADER_ONLY",
			wantHeader: "HEADER_ONLY",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "08. Valid Format (Header Only with Newline)",
			input:      "$HEADER\n",
			wantHeader: "$HEADER",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "09. Valid Format (Header and Whitespace)",
			input:      "$HEADER\n" + "\n\n \n \n",
			wantHeader: "$HEADER",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "10. Valid Format (Whitespace, Header and Body)",
			input:      "  $HEADER  \n  0102  \n\t  0304  \n",
			wantHeader: "$HEADER",
			wantBody:   "01020304",
			wantErr:    false,
		},
		{
			name:        "11. Invalid Input (Empty Input)",
			input:       "",
			wantErr:     true,
			expectedErr: "empty text container",
		},
		{
			name:        "12. Invalid Input (Only Newlines)",
			input:       "\n\n\n",
			wantErr:     true,
			expectedErr: "empty text container",
		},
		{
			name:        "13. Valid Input (Empty Header, Body Only, But Yet Valid As Body Will Be Considered As Header, Will Fail Later When Header Is Checked)",
			input:       "\nbody",
			wantHeader:  "body",
			wantBody:    "",
			wantErr:     false,
			expectedErr: "none",
		},
		{
			name:        "14. Invalid Input (Whitespace Header, Body Only, But Yet Valid As Body Will Be Considered As Header, Will Fail Later When Header Is Checked)",
			input:       " \t\nbody",
			wantHeader:  "body",
			wantBody:    "",
			wantErr:     false,
			expectedErr: "none",
		},
		{
			name:        "15. Invalid Input (Non ASCII header)",
			input:       "£\nbody",
			wantHeader:  "£",
			wantBody:    "body",
			wantErr:     true,
			expectedErr: "invalid characters in header",
		},
		{
			name:        "16. Invalid Input (Non ASCII body)",
			input:       "header\n£",
			wantHeader:  "header",
			wantBody:    "",
			wantErr:     true,
			expectedErr: "invalid characters in body",
		},
		{
			name:        "17. Invalid Input (Control characters header)",
			input:       string([]byte{0x00}) + "header\nbody",
			wantHeader:  "header",
			wantBody:    "body",
			wantErr:     true,
			expectedErr: "invalid characters in header",
		},
		{
			name:        "18. Invalid Input (Control characters body)",
			input:       "header\n\x01",
			wantHeader:  "header",
			wantBody:    "",
			wantErr:     true,
			expectedErr: "invalid characters in body",
		},
		{
			name:       "19. Valid Input (Many newlines)",
			input:      "header\n\n\n\nbody",
			wantHeader: "header",
			wantBody:   "body",
			wantErr:    false,
		},
		{
			name:       "20. Valid Input (Header with internal whitespace)",
			input:      "header with spaces\nbody",
			wantHeader: "header with spaces",
			wantBody:   "body",
			wantErr:    false,
		},
		{
			name:       "21. Valid Input (Body with internal whitespace)",
			input:      "header\nbody with spaces",
			wantHeader: "header",
			wantBody:   "bodywithspaces",
			wantErr:    false,
		},
		{
			name:       "22. Valid Input (Header and body with mixed whitespace)",
			input:      " header \t \n body \t ",
			wantHeader: "header",
			wantBody:   "body",
			wantErr:    false,
		},
		{
			name:       "23. Invalid Input (Header with only whitespace)",
			input:      "    \nbody",
			wantHeader: "body",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:       "24. Valid Input (Body with only whitespace)",
			input:      "header\n    ",
			wantHeader: "header",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name:        "25. Invalid Input (Header and body with only whitespace)",
			input:       "    \n    ",
			wantErr:     true,
			expectedErr: "empty text container",
		},
		{
			name:        "26.Invalid Input (Many invalid Chars)",
			input:       string([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x7F, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F}) + "validHeader\nvalidBody",
			wantHeader:  "validHeader",
			wantBody:    "validBody",
			wantErr:     true,
			expectedErr: "invalid characters in header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, body, err := TextContainerParse(tt.input)

			if (err != nil) != tt.wantErr {
				t.Error(tt.name, `: TextContainerParse() error = `, err, `, wantErr: `, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !strings.Contains(err.Error(), tt.expectedErr) {
					t.Error(tt.name, `: TextContainerParse() error message = `, err.Error(), `, must contain `, tt.expectedErr)
				}
				return // Stop here if we expect an error
			}

			if !reflect.DeepEqual(header, tt.wantHeader) {
				t.Error(tt.name, `: TextContainerParse() header = `, `"`+header+`"`, `, want: `, `"`+tt.wantHeader+`"`)
			}

			if !reflect.DeepEqual(body, tt.wantBody) {
				t.Error(tt.name, `: TextContainerParse() body = `, `"`+body+`"`, `, want:`, `"`+tt.wantBody+`"`)
			}
		})
	}
}
