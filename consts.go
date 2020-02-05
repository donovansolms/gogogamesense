package gogogamesense

// Icon IDs
const (
	// IconDefault will be displayed in the SteelSeries Engine
	IconDefault = 0
	// IconHealth will be displayed in the SteelSeries Engine
	IconHealth = 1
	// IconArmor will be displayed in the SteelSeries Engine
	IconArmor = 2
	// IconAmmo will be displayed in the SteelSeries Engine
	IconAmmo = 3
	// IconMoney will be displayed in the SteelSeries Engine
	IconMoney = 4
	// IconExplosion will be displayed in the SteelSeries Engine
	IconExplosion = 5
	// IconKills will be displayed in the SteelSeries Engine
	IconKills = 6
	// IconHeadshot will be displayed in the SteelSeries Engine
	IconHeadshot = 7
	// IconHelmet will be displayed in the SteelSeries Engine
	IconHelmet = 8
	// IconHunger will be displayed in the SteelSeries Engine
	IconHunger = 10
	// IconAir will be displayed in the SteelSeries Engine
	IconAir = 11
	// IconCompass will be displayed in the SteelSeries Engine
	IconCompass = 12
	// IconPickaxe will be displayed in the SteelSeries Engine
	IconPickaxe = 13
	// IconPotion will be displayed in the SteelSeries Engine
	IconPotion = 14
	// IconClock will be displayed in the SteelSeries Engine
	IconClock = 15
	// IconLightning will be displayed in the SteelSeries Engine
	IconLightning = 16
	// IconBackpack will be displayed in the SteelSeries Engine
	IconBackpack = 17
	// IconAtSymbol will be displayed in the SteelSeries Engine
	IconAtSymbol = 18
	// IconMuted will be displayed in the SteelSeries Engine
	IconMuted = 19
	// IconTalking will be displayed in the SteelSeries Engine
	IconTalking = 20
	// IconConnect will be displayed in the SteelSeries Engine
	IconConnect = 21
	// IconDisconnect will be displayed in the SteelSeries Engine
	IconDisconnect = 22
	// IconMusic will be displayed in the SteelSeries Engine
	IconMusic = 23
	// IconPlay will be displayed in the SteelSeries Engine
	IconPlay = 24
	// IconPause will be displayed in the SteelSeries Engine
	IconPause = 25
)

// Modes
const (
	// ModeCount - As percent, but the number of LEDs illuminated directly correspond
	// to the control value. I.e. if the value is 2, 2 LEDs will be lit. The
	// control value should be between 0 and the size of the zone.
	//
	// *Note*: The count visualization is only enabled for per-key-illuminated
	// devices (e.g. the Apex M800). On other devices, the computed color will
	// be applied to all LEDs, behaving like the color mode.
	ModeCount = "count"
	// ModeColor - All LEDs in the zone are set to the computed color.
	ModeColor = "color"
	// ModePercent - The LEDs are used to create a visualization of the control
	// value (as a percentage) by illuminating them proportionally and leaving
	// the rest black/off. This requires that the control value be in the range
	// of 0-100, inclusive.
	//
	// *Note*: The proportional illumination is only enabled for
	// per-key-illuminated devices
	ModePercent = "percent"
)

// Device types
const (
	// RGB
	// DeviceTypeRGB1Zone - Any connected, supported, single zone RGB device.
	// Covers the Siberia Elite line of headsets, Siberia v3 Prism, Sims 4 line
	// of products, Rival 100, and Rival 110.
	DeviceTypeRGB1Zone = "rgb-1-zone"
	// DeviceTypeRGB2Zone - Any connected, supported, dual zone RGB device.
	// Covers the Rival, Rival 300, Rival 500, and Rival 700 mice lines,
	// Arctis 5, Arctis Pro, QCK Prism Cloth line of mousepads.
	DeviceTypeRGB2Zone = "rgb-2-zone"
	// DeviceTypeRGB3Zone - Any connected, supported, three zone RGB device.
	// Covers the Sensei Wireless mouse and MSI 3 Zone RGB Keyboard.
	DeviceTypeRGB3Zone = "rgb-3-zone"
	// DeviceTypeRGB5Zone - Any connected, supported, five zone RGB device.
	// Covers the Apex 150, Apex 300, and MSI GT72 keyboards.
	DeviceTypeRGB5Zone = "rgb-5-zon"
	// DeviceTypeRGB8Zone - Any connected, supported, eight zone RGB device.
	// Covers the Rival 600 and Rival 650 mice.
	DeviceTypeRGB8Zone = "rgb-8-zone"
	// DeviceTypeRGB12Zone - Any connected, supported, twelve zone RGB device.
	// Covers the QCK Prism mousepad.
	DeviceTypeRGB12Zone = "rgb-12-zone"
	// DeviceTypeRGB17Zone - Any connected, supported, seventeen zone RGB device.
	// Covers the MSI Z270 Gaming Pro Carbon motherboard.
	DeviceTypeRGB17Zone = "rgb-17-zone"
	// DeviceTypeRGB24Zone - Any connected, supported, twenty-four zone RGB device.
	// Covers the MSI Mystic Light.
	DeviceTypeRGB24Zone = "rgb-24-zone"
	// DeviceTypeRGB103Zone - Any connected, supported, one hundred three zone
	// RGB device.
	// Covers the MSI MPG27C and MPG27CQ monitors.
	DeviceTypeRGB103Zone = "rgb-103-zone"
	// DeviceTypeRGBPerKeyZone - Any connected, supported, keyboard with a
	// lighting zone for each key.
	// Covers the Apex M800, Apex M750 and 750 TKL, and MSI RGB Per Key keyboards.
	DeviceTypeRGBPerKeyZone = "rgb-per-key-zones"

	// Tactile
	// DeviceTypeTactile - Any connected, supported device that supports a single
	// zone for tactile feedback.
	// Covers the Rival 500, Rival 700, and Rival 710.
	DeviceTypeTactile = "tactile"

	// Screen
	// DeviceTypeScreened - Any connected, supported device that supports
	// notifications on a single OLED or LCD screen.
	// Covers the Rival 700, Rival 710, Arctis Pro Wireless, and GameDAC.
	// Note: Engine 3.7.0 and later
	DeviceTypeScreened = "screened"
)

// Zones
const (
	// Keyboards
	// DeviceZoneKeyboardFunctionKeys - F1 to F12 (Excludes MSI GE62 and GE72)
	DeviceZoneKeyboardFunctionKeys = "function-keys"
	// DeviceZoneKeyboardMainKeyboard - All keys included in the central area
	// of the keyboard (Excludes MSI GE62 and GE72)
	DeviceZoneKeyboardMainKeyboard = "main-keyboard"
	// DeviceZoneKeyboardKeypad - All keys in the numpad cluster
	// (Excludes MSI GE62 and GE72)
	DeviceZoneKeyboardKeypad = "keypad"
	// DeviceZoneKeyboardNumberKeys - 1-0 on the main keyboard area
	// (Excludes MSI GE62 and GE72)
	DeviceZoneKeyboardNumberKeys = "number-keys"
	// DeviceZoneKeyboardMacroKeys - The macro keys on the left side
	// (Excludes MSI GE62 and GE72)
	DeviceZoneKeyboardMacroKeys = "macro-keys"

	// Mice
	// DeviceZoneMouseWheel - The mousewheel LED
	DeviceZoneMouseWheel = "wheel"
	// DeviceZoneMouseLogo - The logo LED (excludes Sims 4 Mouse)
	DeviceZoneMouseLogo = "logo"
	// DeviceZoneMouseBase - The base station LED (Sensei Wireless only)
	DeviceZoneMouseBase = "base"

	// Headsets
	// DeviceTypeHeadsetEarcups - The LEDs on the earcups
	DeviceTypeHeadsetEarcups = "earcups"

	// Indicator
	// DeviceTypeIndicator - The main LED
	DeviceTypeIndicator = "one"
)
