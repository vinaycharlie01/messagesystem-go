package mon

import (
	"testing"
)

func TestMongoConnect(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := MongoConnect("vinay1", "product")
			if err != nil {
				t.Errorf("MongoConnect() error = %v, wantErr %v", err, tt.name)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("MongoConnect() got = %v, want %v", got, tt.want)
			// }
			// if !reflect.DeepEqual(got1, tt.want1) {
			// 	t.Errorf("MongoConnect() got1 = %v, want %v", got1, tt.want1)
			// }
		})
	}
}
