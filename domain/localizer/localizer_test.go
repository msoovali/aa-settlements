package localizer

import "testing"

func TestLocalizer_Translate(t *testing.T) {
	type fields struct {
		values map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Key found and translated",
			fields: fields{
				values: map[string]interface{}{
					"february": "veebruar",
				},
			},
			args: args{
				key: "february",
			},
			want: "veebruar",
		},
		{
			name: "Key not found returns input",
			fields: fields{
				values: make(map[string]interface{}),
			},
			args: args{
				key: "february",
			},
			want: "february",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Localizer{
				values: tt.fields.values,
			}
			if got := l.Translate(tt.args.key); got != tt.want {
				t.Errorf("Expected %v, received %v", tt.want, got)
			}
		})
	}
}
