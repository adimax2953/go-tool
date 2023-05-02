package geo_ip

import "testing"

func TestCheckBlockIp(t *testing.T) {
	type args struct {
		sourceIp string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "TW",
			args:    args{sourceIp: "150.116.127.222"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "HK",
			args:    args{sourceIp: "175.45.20.138"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "MO",
			args:    args{sourceIp: "122.100.160.253"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "PH",
			args:    args{sourceIp: "115.85.29.130"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "KH",
			args:    args{sourceIp: "124.248.186.156"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckBlockIp(tt.args.sourceIp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckBlockIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckBlockIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
