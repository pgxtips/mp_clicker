package internal

type AppData struct {

}

// initialised on startup (elsewhere)
var globalAppData *AppData

func GetAppData() *AppData {
    if globalAppData == nil {
        panic("GlobalAppData not initialized")
    }
    return globalAppData
}

func SetAppData(a *AppData) {
	globalAppData = a
}
