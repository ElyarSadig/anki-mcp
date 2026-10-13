package anki

import "encoding/json"

// Entity IDs are decimal strings, not JSON numbers.
type EntityId = string
type SessionId = string
type Timestamp = string
type MediaFilename = string
type FieldValues = map[string]string
type Rating = string

type Method struct {
	Available bool   `json:"available"`
	Name      string `json:"name"`
	Reason    any    `json:"reason"`
}

type Description struct {
	ApiMajor    int            `json:"apiMajor"`
	AnkiVersion string         `json:"ankiVersion"`
	Object      map[string]any `json:"object"`
	Methods     []Method       `json:"methods"`
	Limits      map[string]any `json:"limits"`
}

type Command struct {
	Method string `json:"method"`
	Params any    `json:"params"`
}

type ExecutionParams struct {
	Commands []Command `json:"commands"`
}

type ExecutionResponse struct {
	OperationId string  `json:"operationId"`
	Method      string  `json:"method"`
	Session     *string `json:"session"`
}

// Reference selects exactly one existing resource by ID or name.
type Reference struct {
	ID   EntityId `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
}

type SessionParams struct {
	Session SessionId `json:"session,omitempty"`
}

type EntityParams struct {
	ID      EntityId  `json:"id"`
	Session SessionId `json:"session,omitempty"`
}

type EntitiesParams struct {
	IDs     []EntityId `json:"ids"`
	Session SessionId  `json:"session,omitempty"`
}

type SearchParams struct {
	Query   string    `json:"query"`
	Limit   int       `json:"limit,omitempty"`
	Cursor  string    `json:"cursor,omitempty"`
	Sort    *Sort     `json:"sort,omitempty"`
	Session SessionId `json:"session,omitempty"`
}

type Sort struct {
	Column    string `json:"column"`
	Direction string `json:"direction"`
}

type IdPage struct {
	IDs        []EntityId `json:"ids"`
	NextCursor *string    `json:"nextCursor"`
}

type Ack struct {
	Changed bool `json:"changed"`
}
type Count struct {
	ChangedCount int `json:"changedCount"`
}
type EntitySummary struct {
	ID   EntityId `json:"id"`
	Name string   `json:"name"`
}

// Item results preserve the input index, including failures.
type Item[T any] struct {
	Index  int    `json:"index"`
	Result *T     `json:"result,omitempty"`
	Error  *Error `json:"error,omitempty"`
}

type Items[T any] struct {
	Items []Item[T] `json:"items"`
}

type BatchItem struct {
	Index  int             `json:"index"`
	Method string          `json:"method"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

type BatchResult struct {
	Items []BatchItem `json:"items"`
}

type OperationHandle = ExecutionResponse
type OperationParams struct {
	OperationId string `json:"operationId"`
}

type Operation struct {
	OperationId string          `json:"operationId"`
	Method      string          `json:"method"`
	Session     *SessionId      `json:"session"`
	CreatedAt   Timestamp       `json:"createdAt"`
	UpdatedAt   Timestamp       `json:"updatedAt"`
	Cancellable bool            `json:"cancellable"`
	Progress    *Progress       `json:"progress"`
	State       string          `json:"state"`
	Result      json.RawMessage `json:"result,omitempty"`
	Error       *Error          `json:"error,omitempty"`
}

type Progress struct {
	Completed int    `json:"completed"`
	Total     *int   `json:"total"`
	Message   string `json:"message"`
}

type Cancellation struct {
	Requested bool `json:"requested"`
}

type Permission struct {
	Permission    string `json:"permission"`
	RequireApiKey bool   `json:"requireApiKey"`
}

type Profile struct {
	Name string `json:"name"`
}

type SwitchProfileParams struct {
	Name string `json:"name"`
}

type Collection struct {
	Profile   string    `json:"profile"`
	Session   SessionId `json:"session"`
	Scheduler string    `json:"scheduler"`
}

type SyncCollectionParams struct {
	IncludeMedia *bool     `json:"includeMedia,omitempty"`
	Session      SessionId `json:"session,omitempty"`
}

type RefreshCollectionParams struct {
	Target  string    `json:"target,omitempty"`
	Session SessionId `json:"session,omitempty"`
}

type SyncResult struct {
	CollectionComplete bool `json:"collectionComplete"`
	MediaComplete      bool `json:"mediaComplete"`
}

type IntegrityResult struct {
	Problems []string `json:"problems"`
	Repaired bool     `json:"repaired"`
}

type UndoResult struct {
	Label   string `json:"label"`
	Changed bool   `json:"changed"`
}

type Note struct {
	ID         EntityId      `json:"id"`
	NoteType   EntitySummary `json:"noteType"`
	Fields     FieldValues   `json:"fields"`
	Tags       []string      `json:"tags"`
	CardIDs    []EntityId    `json:"cardIds"`
	ModifiedAt Timestamp     `json:"modifiedAt"`
}

type NoteInput struct {
	NoteType             Reference        `json:"noteType"`
	Deck                 Reference        `json:"deck"`
	Fields               FieldValues      `json:"fields"`
	Tags                 []string         `json:"tags,omitempty"`
	Duplicates           *DuplicatePolicy `json:"duplicates,omitempty"`
	RejectEmptyQuestions *bool            `json:"rejectEmptyQuestions,omitempty"`
}

type DuplicatePolicy struct {
	Allow           bool       `json:"allow,omitempty"`
	Scope           string     `json:"scope,omitempty"`
	Deck            *Reference `json:"deck,omitempty"`
	IncludeChildren bool       `json:"includeChildren,omitempty"`
	AcrossNoteTypes bool       `json:"acrossNoteTypes,omitempty"`
}

type ValidateNoteParams struct {
	Note    NoteInput `json:"note"`
	Session SessionId `json:"session,omitempty"`
}

type ValidateNotesParams struct {
	Notes   []NoteInput `json:"notes"`
	Session SessionId   `json:"session,omitempty"`
}

type CreateNoteParams struct {
	NoteInput
	Session SessionId `json:"session,omitempty"`
}

type CreateNotesParams struct {
	Notes   []NoteInput `json:"notes"`
	Atomic  bool        `json:"atomic,omitempty"`
	Session SessionId   `json:"session,omitempty"`
}

type NoteCreated struct {
	NoteID  EntityId   `json:"noteId"`
	CardIDs []EntityId `json:"cardIds"`
}

type NotePatch struct {
	Fields FieldValues `json:"fields,omitempty"`
	Tags   *[]string   `json:"tags,omitempty"`
}

type UpdateNoteParams struct {
	ID      EntityId  `json:"id"`
	Patch   NotePatch `json:"patch"`
	Session SessionId `json:"session,omitempty"`
}

type Validation struct {
	Valid            bool       `json:"valid"`
	Issues           []Issue    `json:"issues"`
	DuplicateNoteIDs []EntityId `json:"duplicateNoteIds"`
}

type Issue struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type PlanConversionParams struct {
	Source  Reference `json:"source"`
	Target  Reference `json:"target"`
	Session SessionId `json:"session,omitempty"`
}

type ConversionMapping struct {
	Target string  `json:"target"`
	Source *string `json:"source"`
}

type ConversionPlan struct {
	Source    EntitySummary       `json:"source"`
	Target    EntitySummary       `json:"target"`
	Fields    []ConversionMapping `json:"fields"`
	Templates []ConversionMapping `json:"templates"`
	Token     string              `json:"token"`
}

type ConvertNotesParams struct {
	IDs       []EntityId          `json:"ids"`
	Target    Reference           `json:"target"`
	Fields    []ConversionMapping `json:"fields"`
	Templates []ConversionMapping `json:"templates"`
	PlanToken string              `json:"planToken"`
	Session   SessionId           `json:"session,omitempty"`
}

type FindEmptyNotesParams struct {
	Criterion string    `json:"criterion"`
	Query     string    `json:"query,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Cursor    string    `json:"cursor,omitempty"`
	Session   SessionId `json:"session,omitempty"`
}

type TagsParams struct {
	NoteIDs []EntityId `json:"noteIds"`
	Tags    []string   `json:"tags"`
	Session SessionId  `json:"session,omitempty"`
}

type ReplaceTagsParams struct {
	From     string      `json:"from"`
	To       string      `json:"to"`
	NoteIDs  *[]EntityId `json:"noteIds,omitempty"`
	AllNotes bool        `json:"allNotes,omitempty"`
	Session  SessionId   `json:"session,omitempty"`
}

type Card struct {
	ID            EntityId        `json:"id"`
	NoteID        EntityId        `json:"noteId"`
	Deck          EntitySummary   `json:"deck"`
	HomeDeck      *EntitySummary  `json:"homeDeck"`
	TemplateIndex int             `json:"templateIndex"`
	Flags         int             `json:"flags"`
	ModifiedAt    Timestamp       `json:"modifiedAt"`
	Scheduling    SchedulingState `json:"scheduling"`
}

type RenderedCard struct {
	CardID       EntityId `json:"cardId"`
	QuestionHtml string   `json:"questionHtml"`
	AnswerHtml   string   `json:"answerHtml"`
	CSS          string   `json:"css"`
}

type MoveCardsParams struct {
	IDs     []EntityId `json:"ids"`
	Deck    Reference  `json:"deck"`
	Session SessionId  `json:"session,omitempty"`
}

type SetSuspendedParams struct {
	IDs       []EntityId `json:"ids"`
	Suspended bool       `json:"suspended"`
	Session   SessionId  `json:"session,omitempty"`
}

type SetFlagsParams struct {
	IDs     []EntityId `json:"ids"`
	Flag    int        `json:"flag"`
	Session SessionId  `json:"session,omitempty"`
}

type CardParams struct {
	CardID  EntityId  `json:"cardId"`
	Session SessionId `json:"session,omitempty"`
}

type DueEligibilityParams struct {
	CardID            EntityId  `json:"cardId"`
	IncludeLearnAhead bool      `json:"includeLearnAhead,omitempty"`
	Session           SessionId `json:"session,omitempty"`
}

type DueEligibility struct {
	CardID     EntityId `json:"cardId"`
	DueNow     bool     `json:"dueNow"`
	Available  bool     `json:"available"`
	LearnAhead bool     `json:"learnAhead"`
}

type SchedulingState struct {
	Type                    string       `json:"type"`
	Queue                   string       `json:"queue"`
	Due                     Due          `json:"due"`
	IntervalDays            int          `json:"intervalDays"`
	LearningIntervalSeconds *int         `json:"learningIntervalSeconds"`
	RemainingSteps          int          `json:"remainingSteps"`
	Repetitions             int          `json:"repetitions"`
	Lapses                  int          `json:"lapses"`
	Sm2EasePermille         int          `json:"sm2EasePermille"`
	FsrsMemory              *MemoryState `json:"fsrsMemory"`
	DesiredRetention        *float64     `json:"desiredRetention"`
}

type Due struct {
	Kind     string     `json:"kind"`
	Position *int       `json:"position,omitempty"`
	At       *Timestamp `json:"at,omitempty"`
	Day      *int       `json:"day,omitempty"`
}

type MemoryState struct {
	Difficulty    float64 `json:"difficulty"`
	StabilityDays float64 `json:"stabilityDays"`
}

type Interval struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Prediction struct {
	Rating   Rating    `json:"rating"`
	Label    string    `json:"label"`
	Interval *Interval `json:"interval"`
}

type Predictions struct {
	CardID     EntityId     `json:"cardId"`
	Source     string       `json:"source"`
	StateToken string       `json:"stateToken"`
	Answers    []Prediction `json:"answers"`
}

type Answer struct {
	CardID     EntityId `json:"cardId"`
	Rating     Rating   `json:"rating"`
	StateToken string   `json:"stateToken"`
	ElapsedMs  int      `json:"elapsedMs,omitempty"`
}

type AnswerParams struct {
	Answer
	Session SessionId `json:"session,omitempty"`
}

type AnswerManyParams struct {
	Answers []Answer  `json:"answers"`
	Session SessionId `json:"session,omitempty"`
}

type AnswerResult struct {
	CardID     EntityId        `json:"cardId"`
	ReviewID   *EntityId       `json:"reviewId"`
	Scheduling SchedulingState `json:"scheduling"`
}

type ResetCardsParams struct {
	CardIDs         []EntityId `json:"cardIds"`
	RestorePosition *bool      `json:"restorePosition,omitempty"`
	ResetCounts     bool       `json:"resetCounts,omitempty"`
	Log             *bool      `json:"log,omitempty"`
	Session         SessionId  `json:"session,omitempty"`
}

type SetDueDateParams struct {
	CardIDs []EntityId `json:"cardIds"`
	Days    string     `json:"days"`
	Session SessionId  `json:"session,omitempty"`
}

type Deck struct {
	ID       EntityId  `json:"id"`
	Name     string    `json:"name"`
	Kind     string    `json:"kind"`
	PresetID *EntityId `json:"presetId"`
}

type DeckParams struct {
	Deck    Reference `json:"deck"`
	Session SessionId `json:"session,omitempty"`
}

type CreateDeckParams struct {
	Name    string    `json:"name"`
	Session SessionId `json:"session,omitempty"`
}

type UpdateDeckParams struct {
	Deck    Reference `json:"deck"`
	Name    string    `json:"name"`
	Session SessionId `json:"session,omitempty"`
}

type DeleteDecksParams struct {
	Decks       []Reference `json:"decks"`
	DeleteCards bool        `json:"deleteCards"`
	Session     SessionId   `json:"session,omitempty"`
}

type DeckCountsParams struct {
	Deck            Reference `json:"deck"`
	IncludeChildren *bool     `json:"includeChildren,omitempty"`
	Session         SessionId `json:"session,omitempty"`
}

type DeckCounts struct {
	DeckID   EntityId `json:"deckId"`
	New      int      `json:"new"`
	Learning int      `json:"learning"`
	Review   int      `json:"review"`
	Total    int      `json:"total"`
}

type DeckPreset struct {
	ID         EntityId       `json:"id"`
	Name       string         `json:"name"`
	Settings   PresetSettings `json:"settings"`
	ModifiedAt Timestamp      `json:"modifiedAt"`
}

type PresetParams struct {
	Preset  Reference `json:"preset"`
	Session SessionId `json:"session,omitempty"`
}

type CreatePresetParams struct {
	Name      string     `json:"name"`
	CloneFrom *Reference `json:"cloneFrom,omitempty"`
	Session   SessionId  `json:"session,omitempty"`
}

type UpdatePresetParams struct {
	Preset         Reference    `json:"preset"`
	Name           string       `json:"name,omitempty"`
	Settings       *PresetPatch `json:"settings,omitempty"`
	RescheduleFsrs bool         `json:"rescheduleFsrs,omitempty"`
	Session        SessionId    `json:"session,omitempty"`
}

type AssignPresetParams struct {
	Decks          []Reference `json:"decks"`
	Preset         Reference   `json:"preset"`
	RescheduleFsrs bool        `json:"rescheduleFsrs,omitempty"`
	Session        SessionId   `json:"session,omitempty"`
}

type DeletePresetParams struct {
	Preset      Reference  `json:"preset"`
	Replacement *Reference `json:"replacement,omitempty"`
	Session     SessionId  `json:"session,omitempty"`
}

type PresetPatch = PresetSettings

// Pointers preserve explicit zero/false settings in partial updates.
type PresetSettings struct {
	LearnStepsMinutes          *[]float64 `json:"learnStepsMinutes,omitempty"`
	RelearnStepsMinutes        *[]float64 `json:"relearnStepsMinutes,omitempty"`
	FsrsParameters4            *[]float64 `json:"fsrsParameters4,omitempty"`
	FsrsParameters5            *[]float64 `json:"fsrsParameters5,omitempty"`
	FsrsParameters6            *[]float64 `json:"fsrsParameters6,omitempty"`
	NewPerDay                  *int       `json:"newPerDay,omitempty"`
	ReviewsPerDay              *int       `json:"reviewsPerDay,omitempty"`
	LeechThreshold             *int       `json:"leechThreshold,omitempty"`
	AnswerTimeLimitSeconds     *int       `json:"answerTimeLimitSeconds,omitempty"`
	MaximumIntervalDays        *int       `json:"maximumIntervalDays,omitempty"`
	MinimumLapseIntervalDays   *int       `json:"minimumLapseIntervalDays,omitempty"`
	GraduatingGoodDays         *int       `json:"graduatingGoodDays,omitempty"`
	GraduatingEasyDays         *int       `json:"graduatingEasyDays,omitempty"`
	InitialEase                *float64   `json:"initialEase,omitempty"`
	EasyMultiplier             *float64   `json:"easyMultiplier,omitempty"`
	HardMultiplier             *float64   `json:"hardMultiplier,omitempty"`
	LapseMultiplier            *float64   `json:"lapseMultiplier,omitempty"`
	IntervalMultiplier         *float64   `json:"intervalMultiplier,omitempty"`
	DesiredRetention           *float64   `json:"desiredRetention,omitempty"`
	HistoricalRetention        *float64   `json:"historicalRetention,omitempty"`
	NewInsertOrder             *string    `json:"newInsertOrder,omitempty"`
	NewGatherOrder             *string    `json:"newGatherOrder,omitempty"`
	NewSortOrder               *string    `json:"newSortOrder,omitempty"`
	ReviewOrder                *string    `json:"reviewOrder,omitempty"`
	NewMix                     *string    `json:"newMix,omitempty"`
	DayLearningMix             *string    `json:"dayLearningMix,omitempty"`
	LeechAction                *string    `json:"leechAction,omitempty"`
	QuestionAction             *string    `json:"questionAction,omitempty"`
	AnswerAction               *string    `json:"answerAction,omitempty"`
	EasyDaysPercentages        *[]float64 `json:"easyDaysPercentages,omitempty"`
	IgnoreReviewsBefore        *string    `json:"ignoreReviewsBefore,omitempty"`
	ParameterSearch            *string    `json:"parameterSearch,omitempty"`
	QuestionDisplaySeconds     *float64   `json:"questionDisplaySeconds,omitempty"`
	AnswerDisplaySeconds       *float64   `json:"answerDisplaySeconds,omitempty"`
	DisableAutoplay            *bool      `json:"disableAutoplay,omitempty"`
	ShowTimer                  *bool      `json:"showTimer,omitempty"`
	StopTimerOnAnswer          *bool      `json:"stopTimerOnAnswer,omitempty"`
	WaitForAudio               *bool      `json:"waitForAudio,omitempty"`
	SkipQuestionOnAnswerReplay *bool      `json:"skipQuestionOnAnswerReplay,omitempty"`
	BuryNew                    *bool      `json:"buryNew,omitempty"`
	BuryReviews                *bool      `json:"buryReviews,omitempty"`
	BuryDayLearning            *bool      `json:"buryDayLearning,omitempty"`
	NewPerDayMinimum           *int       `json:"newPerDayMinimum,omitempty"`
}

type FieldConfig struct {
	Name              string `json:"name"`
	FontFamily        string `json:"fontFamily,omitempty"`
	FontSize          int    `json:"fontSize,omitempty"`
	Description       string `json:"description,omitempty"`
	Sticky            bool   `json:"sticky,omitempty"`
	RTL               bool   `json:"rtl,omitempty"`
	PlainText         bool   `json:"plainText,omitempty"`
	Collapsed         bool   `json:"collapsed,omitempty"`
	ExcludeFromSearch bool   `json:"excludeFromSearch,omitempty"`
}

type FieldPatch struct {
	Name              *string `json:"name,omitempty"`
	FontFamily        *string `json:"fontFamily,omitempty"`
	FontSize          *int    `json:"fontSize,omitempty"`
	Description       *string `json:"description,omitempty"`
	Sticky            *bool   `json:"sticky,omitempty"`
	RTL               *bool   `json:"rtl,omitempty"`
	PlainText         *bool   `json:"plainText,omitempty"`
	Collapsed         *bool   `json:"collapsed,omitempty"`
	ExcludeFromSearch *bool   `json:"excludeFromSearch,omitempty"`
}

type TemplateConfig struct {
	Name                string      `json:"name"`
	QuestionHtml        string      `json:"questionHtml"`
	AnswerHtml          string      `json:"answerHtml"`
	BrowserQuestionHtml string      `json:"browserQuestionHtml,omitempty"`
	BrowserAnswerHtml   string      `json:"browserAnswerHtml,omitempty"`
	Deck                **Reference `json:"deck,omitempty"`
	BrowserFontFamily   string      `json:"browserFontFamily,omitempty"`
	BrowserFontSize     int         `json:"browserFontSize,omitempty"`
}

type TemplatePatch struct {
	Name                *string     `json:"name,omitempty"`
	QuestionHtml        *string     `json:"questionHtml,omitempty"`
	AnswerHtml          *string     `json:"answerHtml,omitempty"`
	BrowserQuestionHtml *string     `json:"browserQuestionHtml,omitempty"`
	BrowserAnswerHtml   *string     `json:"browserAnswerHtml,omitempty"`
	Deck                **Reference `json:"deck,omitempty"`
	BrowserFontFamily   *string     `json:"browserFontFamily,omitempty"`
	BrowserFontSize     *int        `json:"browserFontSize,omitempty"`
}

type CreateNoteTypeParams struct {
	Name        string           `json:"name"`
	Kind        string           `json:"kind,omitempty"`
	Fields      []FieldConfig    `json:"fields"`
	Templates   []TemplateConfig `json:"templates"`
	CSS         string           `json:"css,omitempty"`
	SortField   string           `json:"sortField,omitempty"`
	LatexHeader string           `json:"latexHeader,omitempty"`
	LatexFooter string           `json:"latexFooter,omitempty"`
	LatexSvg    bool             `json:"latexSvg,omitempty"`
	Session     SessionId        `json:"session,omitempty"`
}

type NoteType struct {
	ID          EntityId         `json:"id"`
	Name        string           `json:"name"`
	Kind        string           `json:"kind"`
	Fields      []FieldConfig    `json:"fields"`
	Templates   []TemplateConfig `json:"templates"`
	CSS         string           `json:"css"`
	SortField   string           `json:"sortField"`
	LatexHeader string           `json:"latexHeader"`
	LatexFooter string           `json:"latexFooter"`
	LatexSvg    bool             `json:"latexSvg"`
	ModifiedAt  Timestamp        `json:"modifiedAt"`
}

type NoteTypeParams struct {
	NoteType Reference `json:"noteType"`
	Session  SessionId `json:"session,omitempty"`
}

type NoteTypePatch struct {
	Name        *string `json:"name,omitempty"`
	CSS         *string `json:"css,omitempty"`
	SortField   *string `json:"sortField,omitempty"`
	LatexHeader *string `json:"latexHeader,omitempty"`
	LatexFooter *string `json:"latexFooter,omitempty"`
	LatexSvg    *bool   `json:"latexSvg,omitempty"`
}

type UpdateNoteTypeParams struct {
	NoteType Reference     `json:"noteType"`
	Patch    NoteTypePatch `json:"patch"`
	Session  SessionId     `json:"session,omitempty"`
}

// Patch is a FieldPatch for updateField and a TemplatePatch for updateTemplate.
type StructureEdit struct {
	Op       string          `json:"op"`
	Name     string          `json:"name,omitempty"`
	Index    *int            `json:"index,omitempty"`
	Field    *FieldConfig    `json:"field,omitempty"`
	Template *TemplateConfig `json:"template,omitempty"`
	Patch    any             `json:"patch,omitempty"`
}

type EditStructureParams struct {
	NoteType Reference       `json:"noteType"`
	Edits    []StructureEdit `json:"edits"`
	Session  SessionId       `json:"session,omitempty"`
}

type TemplateUsage struct {
	TemplateName   string   `json:"templateName"`
	QuestionFields []string `json:"questionFields"`
	AnswerFields   []string `json:"answerFields"`
}

type ReplaceTextParams struct {
	Find         string       `json:"find"`
	Replace      string       `json:"replace"`
	NoteTypes    *[]Reference `json:"noteTypes,omitempty"`
	AllNoteTypes bool         `json:"allNoteTypes,omitempty"`
	Targets      []string     `json:"targets,omitempty"`
	Session      SessionId    `json:"session,omitempty"`
}

type DeleteNoteTypesParams struct {
	NoteTypes []Reference `json:"noteTypes"`
	Session   SessionId   `json:"session,omitempty"`
}

// MediaSource requires exactly one source. Data may contain an empty base64 string.
type MediaSource struct {
	Data *string `json:"data,omitempty"`
	Path string  `json:"path,omitempty"`
	URL  string  `json:"url,omitempty"`
}

type MediaSkip struct {
	Algorithm string `json:"algorithm"`
	Digest    string `json:"digest"`
}

type StoreMediaParams struct {
	Filename   MediaFilename `json:"filename"`
	Source     MediaSource   `json:"source"`
	Collision  string        `json:"collision,omitempty"`
	SkipIfHash *MediaSkip    `json:"skipIfHash,omitempty"`
	Session    SessionId     `json:"session,omitempty"`
}

type MediaParams struct {
	Filename MediaFilename `json:"filename"`
	Session  SessionId     `json:"session,omitempty"`
}

type ListMediaParams struct {
	Pattern MediaFilename `json:"pattern,omitempty"`
	Session SessionId     `json:"session,omitempty"`
}

type MediaFile struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
}

type MediaStored struct {
	Filename *string `json:"filename"`
	Status   string  `json:"status"`
	SHA256   string  `json:"sha256"`
}

type MediaDirectory struct {
	Path string `json:"path"`
}

type ListReviewsParams struct {
	CardIDs         *[]EntityId `json:"cardIds,omitempty"`
	Deck            *Reference  `json:"deck,omitempty"`
	IncludeChildren bool        `json:"includeChildren,omitempty"`
	Membership      string      `json:"membership,omitempty"`
	AfterID         EntityId    `json:"afterId,omitempty"`
	IncludeManual   *bool       `json:"includeManual,omitempty"`
	Limit           int         `json:"limit,omitempty"`
	Cursor          string      `json:"cursor,omitempty"`
	Session         SessionId   `json:"session,omitempty"`
}

type LatestReviewParams struct {
	Deck       *Reference  `json:"deck,omitempty"`
	CardIDs    *[]EntityId `json:"cardIds,omitempty"`
	Membership string      `json:"membership,omitempty"`
	Session    SessionId   `json:"session,omitempty"`
}

type LatestReviewId struct {
	ID *EntityId `json:"id"`
}

type ReviewImport struct {
	ID               EntityId       `json:"id"`
	CardID           EntityId       `json:"cardId"`
	Rating           *Rating        `json:"rating"`
	Interval         Interval       `json:"interval"`
	PreviousInterval Interval       `json:"previousInterval"`
	Factor           map[string]any `json:"factor"`
	ElapsedMs        int            `json:"elapsedMs"`
	Kind             string         `json:"kind"`
}

type Review struct {
	ReviewImport
	AnsweredAt   Timestamp `json:"answeredAt"`
	SyncSequence int       `json:"syncSequence"`
}

type ReviewPage struct {
	Items      []Review `json:"items"`
	NextCursor *string  `json:"nextCursor"`
}

type ReviewCount struct {
	Count     int       `json:"count"`
	Day       string    `json:"day"`
	NextDayAt Timestamp `json:"nextDayAt"`
}

type ReviewDay struct {
	Day   string `json:"day"`
	Count int    `json:"count"`
}

type StatisticsParams struct {
	Deck    *Reference `json:"deck,omitempty"`
	Session SessionId  `json:"session,omitempty"`
}

type StatisticsSummary struct {
	CardCount       int   `json:"cardCount"`
	NoteCount       int   `json:"noteCount"`
	ReviewCount     int   `json:"reviewCount"`
	ReviewTimeMs    int64 `json:"reviewTimeMs"`
	NewCount        int   `json:"newCount"`
	LearningCount   int   `json:"learningCount"`
	ReviewCardCount int   `json:"reviewCardCount"`
	SuspendedCount  int   `json:"suspendedCount"`
}

type StatisticsHtml struct {
	Html   string `json:"html"`
	Format string `json:"format"`
}

type ImportPackageParams struct {
	Path    string    `json:"path"`
	Session SessionId `json:"session,omitempty"`
}

type ExportPackageParams struct {
	Deck              Reference `json:"deck"`
	Path              string    `json:"path"`
	IncludeScheduling bool      `json:"includeScheduling,omitempty"`
	Session           SessionId `json:"session,omitempty"`
}

type ImportResult struct {
	Completed bool `json:"completed"`
}

type ExportResult struct {
	Path string `json:"path"`
}

type OpenBrowserParams struct {
	Query   *string   `json:"query,omitempty"`
	Sort    *Sort     `json:"sort,omitempty"`
	Session SessionId `json:"session,omitempty"`
}

type BrowserState struct {
	Open            bool       `json:"open"`
	SelectedNoteIDs []EntityId `json:"selectedNoteIds"`
	SelectedCardIDs []EntityId `json:"selectedCardIds"`
}

type SelectBrowserParams struct {
	CardIDs *[]EntityId `json:"cardIds,omitempty"`
	NoteIDs *[]EntityId `json:"noteIds,omitempty"`
	Session SessionId   `json:"session,omitempty"`
}

type OpenNoteEditorParams struct {
	NoteID  EntityId  `json:"noteId"`
	Session SessionId `json:"session,omitempty"`
}

type DraftInput struct {
	NoteType Reference   `json:"noteType"`
	Deck     Reference   `json:"deck"`
	Fields   FieldValues `json:"fields,omitempty"`
	Tags     []string    `json:"tags,omitempty"`
}

type OpenAddEditorParams struct {
	Note    *DraftInput `json:"note,omitempty"`
	Session SessionId   `json:"session,omitempty"`
}

type EditorPatch struct {
	Deck     *Reference  `json:"deck,omitempty"`
	NoteType *Reference  `json:"noteType,omitempty"`
	Fields   FieldValues `json:"fields,omitempty"`
	Tags     *[]string   `json:"tags,omitempty"`
}

type UpdateDraftParams struct {
	EditorID string      `json:"editorId"`
	Patch    EditorPatch `json:"patch"`
	Mode     string      `json:"mode,omitempty"`
	Session  SessionId   `json:"session,omitempty"`
}

type EditorState struct {
	ID       string         `json:"id"`
	Kind     string         `json:"kind"`
	Variant  string         `json:"variant"`
	NoteID   *EntityId      `json:"noteId"`
	Deck     *EntitySummary `json:"deck"`
	NoteType EntitySummary  `json:"noteType"`
	Fields   FieldValues    `json:"fields"`
	Tags     []string       `json:"tags"`
	Dirty    bool           `json:"dirty"`
	Session  SessionId      `json:"session"`
}

type ReviewerState struct {
	Active     bool          `json:"active"`
	Card       *Card         `json:"card"`
	Rendered   *RenderedCard `json:"rendered"`
	Side       string        `json:"side"`
	StateToken *string       `json:"stateToken"`
	Answers    []Prediction  `json:"answers"`
	Source     string        `json:"source"`
}

type ReviewerTokenParams struct {
	StateToken string    `json:"stateToken"`
	Session    SessionId `json:"session,omitempty"`
}

type ReviewerAnswerParams struct {
	StateToken string    `json:"stateToken"`
	Rating     Rating    `json:"rating"`
	Session    SessionId `json:"session,omitempty"`
}

type NavigateParams struct {
	View    string     `json:"view"`
	Deck    *Reference `json:"deck,omitempty"`
	Session SessionId  `json:"session,omitempty"`
}

type ImportFileParams struct {
	Path    string    `json:"path,omitempty"`
	Session SessionId `json:"session,omitempty"`
}

type ExitResponse struct {
	Accepted bool `json:"accepted"`
}

// A nil outer pointer omits a nullable patch; a pointer to nil sends JSON null.
type RawCardPatch struct {
	IntervalDays    *int          `json:"intervalDays,omitempty"`
	Sm2EasePermille *int          `json:"sm2EasePermille,omitempty"`
	Repetitions     *int          `json:"repetitions,omitempty"`
	Lapses          *int          `json:"lapses,omitempty"`
	RemainingSteps  *int          `json:"remainingSteps,omitempty"`
	FsrsMemory      **MemoryState `json:"fsrsMemory,omitempty"`
	CustomData      *string       `json:"customData,omitempty"`
}

type UpdateCardParams struct {
	CardID  EntityId     `json:"cardId"`
	Patch   RawCardPatch `json:"patch"`
	Session SessionId    `json:"session,omitempty"`
}

type ImportReviewsParams struct {
	Reviews []ReviewImport `json:"reviews"`
	Session SessionId      `json:"session,omitempty"`
}
