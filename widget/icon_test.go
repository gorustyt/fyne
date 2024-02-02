package widget_test

import (
	"testing"

	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/layout"
	"github.com/gorustyt/fyne/v2/test"
	"github.com/gorustyt/fyne/v2/theme"
	"github.com/gorustyt/fyne/v2/widget"
)

func TestIcon_Layout(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		resource fyne.Resource
	}{
		"empty": {},
		"resource": {
			resource: theme.CancelIcon(),
		},
	} {
		t.Run(name, func(t *testing.T) {
			icon := &widget.Icon{
				Resource: tt.resource,
			}

			window := test.NewWindow(&fyne.Container{Layout: layout.NewCenterLayout(), Objects: []fyne.CanvasObject{icon}})
			window.Resize(icon.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, "icon/layout_"+name+".xml", window.Canvas())

			window.Close()
		})
	}
}
