package version

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/kubescape/kubescape/v3/core/core"

	"github.com/kubescape/backend/pkg/versioncheck"
	"github.com/stretchr/testify/assert"
)

func TestGetVersionCmd(t *testing.T) {
	tests := []struct {
		name        string
		buildNumber string
		want        string
	}{
		{
			name:        "Undefined Build Number",
			buildNumber: "unknown",
			want:        "Your current version is: unknown\n",
		},
		{
			name:        "Defined Build Number: v3.0.1",
			buildNumber: "v3.0.1",
			want:        "Your current version is: v3.0.1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versioncheck.BuildNumber = tt.buildNumber

			ks := core.NewKubescape(context.TODO())
			if cmd := GetVersionCmd(ks); cmd != nil {
				buf := bytes.NewBufferString("")
				cmd.SetOut(buf)
				cmd.Execute()
				out, err := io.ReadAll(buf)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.want, string(out))
			}
		})
	}
}
