package vsc

type VsCodeTheme struct {
	Name                 string   `json:"name"`
	Type                 string   `json:"type"`
	SemanticHighlighting bool     `json:"semanticHighlighting"`
	Colors               VSColors `json:"colors"`
}

type VSColors struct {
	// Terminal Colors
	VSTerminal
	ContrastBorder                               string      `json:"contrastBorder"`
	FocusBorder                                  string      `json:"focusBorder"`
	Foreground                                   string      `json:"foreground"`
	WidgetShadow                                 string      `json:"widget.shadow"`
	SelectionBackground                          string      `json:"selection.background"`
	ErrorForeground                              string      `json:"errorForeground"`
	ButtonBackground                             string      `json:"button.background"`
	ButtonForeground                             string      `json:"button.foreground"`
	ButtonHoverBackground                        string      `json:"button.hoverBackground"`
	DropdownBackground                           string      `json:"dropdown.background"`
	DropdownBorder                               string      `json:"dropdown.border"`
	DropdownForeground                           string      `json:"dropdown.foreground"`
	InputBackground                              string      `json:"input.background"`
	InputBorder                                  string      `json:"input.border"`
	InputForeground                              string      `json:"input.foreground"`
	InputPlaceholderForeground                   string      `json:"input.placeholderForeground"`
	InputOptionActiveBorder                      string      `json:"inputOption.activeBorder"`
	PunctuationDefinitionGenericBeginHTML        string      `json:"punctuation.definition.generic.begin.html"`
	InputValidationErrorBackground               string      `json:"inputValidation.errorBackground"`
	InputValidationErrorBorder                   string      `json:"inputValidation.errorBorder"`
	InputValidationInfoBackground                string      `json:"inputValidation.infoBackground"`
	InputValidationInfoBorder                    string      `json:"inputValidation.infoBorder"`
	InputValidationWarningBackground             string      `json:"inputValidation.warningBackground"`
	InputValidationWarningBorder                 string      `json:"inputValidation.warningBorder"`
	ScrollbarShadow                              string      `json:"scrollbar.shadow"`
	ScrollbarSliderActiveBackground              string      `json:"scrollbarSlider.activeBackground"`
	ScrollbarSliderBackground                    string      `json:"scrollbarSlider.background"`
	ScrollbarSliderHoverBackground               string      `json:"scrollbarSlider.hoverBackground"`
	BadgeBackground                              string      `json:"badge.background"`
	BadgeForeground                              string      `json:"badge.foreground"`
	ProgressBackground                           string      `json:"progress.background"`
	BreadcrumbForeground                         string      `json:"breadcrumb.foreground"`
	BreadcrumbFocusForeground                    string      `json:"breadcrumb.focusForeground"`
	BreadcrumbActiveSelectionForeground          string      `json:"breadcrumb.activeSelectionForeground"`
	BreadcrumbPickerBackground                   string      `json:"breadcrumbPicker.background"`
	ListActiveSelectionBackground                string      `json:"list.activeSelectionBackground"`
	ListActiveSelectionForeground                string      `json:"list.activeSelectionForeground"`
	ListInvalidItemForeground                    string      `json:"list.invalidItemForeground"`
	ListDropBackground                           string      `json:"list.dropBackground"`
	ListFocusBackground                          string      `json:"list.focusBackground"`
	ListFocusForeground                          string      `json:"list.focusForeground"`
	ListHighlightForeground                      string      `json:"list.highlightForeground"`
	ListHoverBackground                          string      `json:"list.hoverBackground"`
	ListHoverForeground                          string      `json:"list.hoverForeground"`
	ListInactiveSelectionBackground              string      `json:"list.inactiveSelectionBackground"`
	ListInactiveSelectionForeground              string      `json:"list.inactiveSelectionForeground"`
	ActivityBarBackground                        string      `json:"activityBar.background"`
	ActivityBarDropBackground                    string      `json:"activityBar.dropBackground"`
	ActivityBarForeground                        string      `json:"activityBar.foreground"`
	ActivityBarBorder                            string      `json:"activityBar.border"`
	ActivityBarBadgeBackground                   string      `json:"activityBarBadge.background"`
	ActivityBarBadgeForeground                   string      `json:"activityBarBadge.foreground"`
	SideBarBackground                            string      `json:"sideBar.background"`
	SideBarForeground                            string      `json:"sideBar.foreground"`
	SideBarBorder                                string      `json:"sideBar.border"`
	SideBarTitleForeground                       string      `json:"sideBarTitle.foreground"`
	SideBarSectionHeaderBackground               string      `json:"sideBarSectionHeader.background"`
	SideBarSectionHeaderForeground               string      `json:"sideBarSectionHeader.foreground"`
	EditorGroupEmptyBackground                   string      `json:"editorGroup.emptyBackground"`
	EditorGroupBorder                            string      `json:"editorGroup.border"`
	EditorGroupDropBackground                    string      `json:"editorGroup.dropBackground"`
	EditorGroupHeaderNoTabsBackground            string      `json:"editorGroupHeader.noTabsBackground"`
	EditorGroupHeaderTabsBackground              string      `json:"editorGroupHeader.tabsBackground"`
	EditorGroupHeaderTabsBorder                  string      `json:"editorGroupHeader.tabsBorder"`
	TabActiveBackground                          string      `json:"tab.activeBackground"`
	TabActiveForeground                          string      `json:"tab.activeForeground"`
	TabBorder                                    string      `json:"tab.border"`
	TabActiveBorder                              string      `json:"tab.activeBorder"`
	TabUnfocusedActiveBorder                     string      `json:"tab.unfocusedActiveBorder"`
	TabInactiveBackground                        string      `json:"tab.inactiveBackground"`
	TabInactiveForeground                        string      `json:"tab.inactiveForeground"`
	TabUnfocusedActiveForeground                 string      `json:"tab.unfocusedActiveForeground"`
	TabUnfocusedInactiveForeground               string      `json:"tab.unfocusedInactiveForeground"`
	EditorBackground                             string      `json:"editor.background"`
	EditorForeground                             string      `json:"editor.foreground"`
	EditorLineNumberForeground                   string      `json:"editorLineNumber.foreground"`
	EditorLineNumberActiveForeground             string      `json:"editorLineNumber.activeForeground"`
	EditorCursorForeground                       string      `json:"editorCursor.foreground"`
	EditorSelectionBackground                    string      `json:"editor.selectionBackground"`
	EditorSelectionHighlightBackground           string      `json:"editor.selectionHighlightBackground"`
	EditorInactiveSelectionBackground            string      `json:"editor.inactiveSelectionBackground"`
	EditorWordHighlightBackground                string      `json:"editor.wordHighlightBackground"`
	EditorWordHighlightStrongBackground          string      `json:"editor.wordHighlightStrongBackground"`
	EditorFindMatchBackground                    string      `json:"editor.findMatchBackground"`
	EditorFindMatchHighlightBackground           string      `json:"editor.findMatchHighlightBackground"`
	EditorFindRangeHighlightBackground           interface{} `json:"editor.findRangeHighlightBackground"`
	EditorHoverHighlightBackground               string      `json:"editor.hoverHighlightBackground"`
	EditorLineHighlightBackground                string      `json:"editor.lineHighlightBackground"`
	EditorLineHighlightBorder                    interface{} `json:"editor.lineHighlightBorder"`
	EditorLinkActiveForeground                   interface{} `json:"editorLink.activeForeground"`
	EditorRangeHighlightBackground               string      `json:"editor.rangeHighlightBackground"`
	EditorWhitespaceForeground                   interface{} `json:"editorWhitespace.foreground"`
	EditorIndentGuideBackground                  string      `json:"editorIndentGuide.background"`
	EditorIndentGuideActiveBackground            string      `json:"editorIndentGuide.activeBackground"`
	EditorRulerForeground                        string      `json:"editorRuler.foreground"`
	EditorCodeLensForeground                     string      `json:"editorCodeLens.foreground"`
	EditorBracketMatchBackground                 string      `json:"editorBracketMatch.background"`
	EditorBracketMatchBorder                     interface{} `json:"editorBracketMatch.border"`
	EditorOverviewRulerCurrentContentForeground  string      `json:"editorOverviewRuler.currentContentForeground"`
	EditorOverviewRulerIncomingContentForeground string      `json:"editorOverviewRuler.incomingContentForeground"`
	EditorOverviewRulerCommonContentForeground   string      `json:"editorOverviewRuler.commonContentForeground"`
	EditorErrorForeground                        string      `json:"editorError.foreground"`
	EditorErrorBorder                            interface{} `json:"editorError.border"`
	EditorWarningForeground                      string      `json:"editorWarning.foreground"`
	EditorWarningBorder                          interface{} `json:"editorWarning.border"`
	EditorGutterBackground                       string      `json:"editorGutter.background"`
	EditorGutterModifiedBackground               string      `json:"editorGutter.modifiedBackground"`
	EditorGutterAddedBackground                  string      `json:"editorGutter.addedBackground"`
	EditorGutterDeletedBackground                string      `json:"editorGutter.deletedBackground"`
	DiffEditorInsertedTextBackground             string      `json:"diffEditor.insertedTextBackground"`
	DiffEditorInsertedTextBorder                 string      `json:"diffEditor.insertedTextBorder"`
	DiffEditorRemovedTextBackground              string      `json:"diffEditor.removedTextBackground"`
	DiffEditorRemovedTextBorder                  string      `json:"diffEditor.removedTextBorder"`
	EditorWidgetBackground                       string      `json:"editorWidget.background"`
	EditorWidgetBorder                           string      `json:"editorWidget.border"`
	EditorSuggestWidgetBackground                string      `json:"editorSuggestWidget.background"`
	EditorSuggestWidgetBorder                    string      `json:"editorSuggestWidget.border"`
	EditorSuggestWidgetForeground                string      `json:"editorSuggestWidget.foreground"`
	EditorSuggestWidgetHighlightForeground       string      `json:"editorSuggestWidget.highlightForeground"`
	EditorSuggestWidgetSelectedBackground        string      `json:"editorSuggestWidget.selectedBackground"`
	EditorHoverWidgetBackground                  string      `json:"editorHoverWidget.background"`
	EditorHoverWidgetBorder                      string      `json:"editorHoverWidget.border"`
	DebugExceptionWidgetBackground               string      `json:"debugExceptionWidget.background"`
	DebugExceptionWidgetBorder                   string      `json:"debugExceptionWidget.border"`
	EditorMarkerNavigationBackground             string      `json:"editorMarkerNavigation.background"`
	EditorMarkerNavigationErrorBackground        string      `json:"editorMarkerNavigationError.background"`
	EditorMarkerNavigationWarningBackground      string      `json:"editorMarkerNavigationWarning.background"`
	PeekViewBorder                               string      `json:"peekView.border"`
	PeekViewEditorBackground                     string      `json:"peekViewEditor.background"`
	PeekViewEditorMatchHighlightBackground       string      `json:"peekViewEditor.matchHighlightBackground"`
	PeekViewResultBackground                     string      `json:"peekViewResult.background"`
	PeekViewResultFileForeground                 string      `json:"peekViewResult.fileForeground"`
	PeekViewResultLineForeground                 string      `json:"peekViewResult.lineForeground"`
	PeekViewResultMatchHighlightBackground       string      `json:"peekViewResult.matchHighlightBackground"`
	PeekViewResultSelectionBackground            string      `json:"peekViewResult.selectionBackground"`
	PeekViewResultSelectionForeground            string      `json:"peekViewResult.selectionForeground"`
	PeekViewTitleBackground                      string      `json:"peekViewTitle.background"`
	PeekViewTitleDescriptionForeground           string      `json:"peekViewTitleDescription.foreground"`
	PeekViewTitleLabelForeground                 string      `json:"peekViewTitleLabel.foreground"`
	MergeCurrentHeaderBackground                 string      `json:"merge.currentHeaderBackground"`
	MergeCurrentContentBackground                interface{} `json:"merge.currentContentBackground"`
	MergeIncomingHeaderBackground                string      `json:"merge.incomingHeaderBackground"`
	MergeIncomingContentBackground               interface{} `json:"merge.incomingContentBackground"`
	MergeBorder                                  interface{} `json:"merge.border"`
	PanelBackground                              string      `json:"panel.background"`
	PanelBorder                                  string      `json:"panel.border"`
	PanelTitleActiveBorder                       string      `json:"panelTitle.activeBorder"`
	PanelTitleActiveForeground                   string      `json:"panelTitle.activeForeground"`
	PanelTitleInactiveForeground                 string      `json:"panelTitle.inactiveForeground"`
	StatusBarBackground                          string      `json:"statusBar.background"`
	StatusBarForeground                          string      `json:"statusBar.foreground"`
	StatusBarBorder                              string      `json:"statusBar.border"`
	StatusBarDebuggingBackground                 string      `json:"statusBar.debuggingBackground"`
	StatusBarDebuggingForeground                 interface{} `json:"statusBar.debuggingForeground"`
	StatusBarDebuggingBorder                     string      `json:"statusBar.debuggingBorder"`
	StatusBarNoFolderForeground                  interface{} `json:"statusBar.noFolderForeground"`
	StatusBarNoFolderBackground                  string      `json:"statusBar.noFolderBackground"`
	StatusBarNoFolderBorder                      string      `json:"statusBar.noFolderBorder"`
	StatusBarItemActiveBackground                string      `json:"statusBarItem.activeBackground"`
	StatusBarItemHoverBackground                 string      `json:"statusBarItem.hoverBackground"`
	StatusBarItemProminentBackground             string      `json:"statusBarItem.prominentBackground"`
	StatusBarItemProminentHoverBackground        string      `json:"statusBarItem.prominentHoverBackground"`
	TitleBarActiveBackground                     string      `json:"titleBar.activeBackground"`
	TitleBarActiveForeground                     string      `json:"titleBar.activeForeground"`
	TitleBarInactiveBackground                   string      `json:"titleBar.inactiveBackground"`
	TitleBarInactiveForeground                   interface{} `json:"titleBar.inactiveForeground"`
	NotificationsBackground                      string      `json:"notifications.background"`
	NotificationsBorder                          string      `json:"notifications.border"`
	NotificationCenterBorder                     string      `json:"notificationCenter.border"`
	NotificationToastBorder                      string      `json:"notificationToast.border"`
	NotificationsForeground                      string      `json:"notifications.foreground"`
	NotificationLinkForeground                   string      `json:"notificationLink.foreground"`
	ExtensionButtonProminentForeground           string      `json:"extensionButton.prominentForeground"`
	ExtensionButtonProminentBackground           string      `json:"extensionButton.prominentBackground"`
	ExtensionButtonProminentHoverBackground      string      `json:"extensionButton.prominentHoverBackground"`
	PickerGroupForeground                        string      `json:"pickerGroup.foreground"`
	PickerGroupBorder                            string      `json:"pickerGroup.border"`

	TextCodeBlockBackground                    string `json:"textCodeBlock.background"`
	DebugToolBarBackground                     string `json:"debugToolBar.background"`
	WelcomePageButtonBackground                string `json:"welcomePage.buttonBackground"`
	WelcomePageButtonHoverBackground           string `json:"welcomePage.buttonHoverBackground"`
	WalkThroughEmbeddedEditorBackground        string `json:"walkThrough.embeddedEditorBackground"`
	GitDecorationModifiedResourceForeground    string `json:"gitDecoration.modifiedResourceForeground"`
	GitDecorationDeletedResourceForeground     string `json:"gitDecoration.deletedResourceForeground"`
	GitDecorationUntrackedResourceForeground   string `json:"gitDecoration.untrackedResourceForeground"`
	GitDecorationIgnoredResourceForeground     string `json:"gitDecoration.ignoredResourceForeground"`
	GitDecorationConflictingResourceForeground string `json:"gitDecoration.conflictingResourceForeground"`
	SourceElm                                  string `json:"source.elm"`
	StringQuotedSingleJs                       string `json:"string.quoted.single.js"`
	MetaObjectliteralJs                        string `json:"meta.objectliteral.js"`
}

type VSTerminal struct {
	TerminalAnsiWhite           string `json:"terminal.ansiWhite"`
	TerminalAnsiBlack           string `json:"terminal.ansiBlack"`
	TerminalAnsiBlue            string `json:"terminal.ansiBlue"`
	TerminalAnsiCyan            string `json:"terminal.ansiCyan"`
	TerminalAnsiGreen           string `json:"terminal.ansiGreen"`
	TerminalAnsiMagenta         string `json:"terminal.ansiMagenta"`
	TerminalAnsiRed             string `json:"terminal.ansiRed"`
	TerminalAnsiYellow          string `json:"terminal.ansiYellow"`
	TerminalAnsiBrightWhite     string `json:"terminal.ansiBrightWhite"`
	TerminalAnsiBrightBlack     string `json:"terminal.ansiBrightBlack"`
	TerminalAnsiBrightBlue      string `json:"terminal.ansiBrightBlue"`
	TerminalAnsiBrightCyan      string `json:"terminal.ansiBrightCyan"`
	TerminalAnsiBrightGreen     string `json:"terminal.ansiBrightGreen"`
	TerminalAnsiBrightMagenta   string `json:"terminal.ansiBrightMagenta"`
	TerminalAnsiBrightRed       string `json:"terminal.ansiBrightRed"`
	TerminalAnsiBrightYellow    string `json:"terminal.ansiBrightYellow"`
	TerminalSelectionBackground string `json:"terminal.selectionBackground"`
	TerminalCursorBackground    string `json:"terminalCursor.background"`
}

type PackageJson struct {
	Contributes struct {
		Themes []struct {
			Label   string `json:"label"`
			UITheme string `json:"uiTheme"`
			Path    string `json:"path"`
		} `json:"themes"`
	} `json:"contributes"`
}

// TODO: Get list of all common vscode textmate grammars and use
//Input
type TokenColors struct {
	Name     string    `json:"name"`
	Scope    []string  `json:"scope"`
	Settings Xsettings `json:"settings"`
}

type Xsettings struct {
	Foreground string `json:"foreground"`
	Background string `json:"background"`
	FontStyle  string `json:"fontStyle"`
}

// TODO: Parse this properly
var Common = []string{
	"comment",
	"constant",
	/**/ "constant.numeric",
	/**/ "constant.character",
	/**/ "constant.language",

	"invalid",

	"keyword",
	/**/ "keyword.control",
	/**/ "keyword.operator",
	/**/ "keyword.other",

	"markup",
	"meta",

	"storage",          // Things relating to storage
	"storage.type",     //   eg: class, function, int, var
	"storage.modifier", //   eg: static final abstract

	"string",
	"string.quoted", //
	"string.quoted.{single, double, triple}",
	"string.unquoted",
	"string.interpolated",
	"string.regexp",

	"variable",
	"variable.parameter",
	"variable.language",
	"variable.other",
	"meta",
	"entity",

	"markup",
}
