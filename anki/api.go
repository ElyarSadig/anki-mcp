package anki

import "context"

var emptyPrams = map[string]struct{}{}

func (c *AnkiClient) Describe(ctx context.Context) (Description, error) {
	const method = "api.describe"

	var resp Description
	err := c.call(ctx, method, emptyPrams, &resp)
	if err != nil {
		return Description{}, err
	}

	return resp, nil
}

func (c *AnkiClient) Execute(ctx context.Context, params ExecutionParams) (ExecutionResponse, error) {
	const method = "api.execute"

	var resp ExecutionResponse
	err := c.call(ctx, method, params, &resp)
	if err != nil {
		return ExecutionResponse{}, err
	}

	return resp, nil
}

func (c *AnkiClient) RequestPermission(ctx context.Context) (Permission, error) {
	const method = "access.requestPermission"
	var resp Permission
	if err := c.call(ctx, method, emptyPrams, &resp); err != nil {
		return Permission{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListProfiles(ctx context.Context) ([]Profile, error) {
	const method = "profiles.list"
	var resp []Profile
	if err := c.call(ctx, method, emptyPrams, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) GetActiveProfile(ctx context.Context) (*Profile, error) {
	const method = "profiles.getActive"
	var resp *Profile
	if err := c.call(ctx, method, emptyPrams, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) SwitchProfile(ctx context.Context, params SwitchProfileParams) (OperationHandle, error) {
	const method = "profiles.switch"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetCollection(ctx context.Context, params SessionParams) (Collection, error) {
	const method = "collection.get"
	var resp Collection
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Collection{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SyncCollection(ctx context.Context, params SyncCollectionParams) (OperationHandle, error) {
	const method = "collection.sync"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) RefreshCollection(ctx context.Context, params RefreshCollectionParams) (Ack, error) {
	const method = "collection.refresh"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CheckIntegrity(ctx context.Context, params SessionParams) (OperationHandle, error) {
	const method = "collection.checkIntegrity"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) Undo(ctx context.Context, params SessionParams) (OperationHandle, error) {
	const method = "collection.undo"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetNote(ctx context.Context, params EntityParams) (Note, error) {
	const method = "notes.get"
	var resp Note
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Note{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetNotes(ctx context.Context, params EntitiesParams) (Items[Note], error) {
	const method = "notes.getMany"
	var resp Items[Note]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Note]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SearchNotes(ctx context.Context, params SearchParams) (IdPage, error) {
	const method = "notes.search"
	var resp IdPage
	if err := c.call(ctx, method, params, &resp); err != nil {
		return IdPage{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ValidateNote(ctx context.Context, params ValidateNoteParams) (Validation, error) {
	const method = "notes.validate"
	var resp Validation
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Validation{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ValidateNotes(ctx context.Context, params ValidateNotesParams) (Items[Validation], error) {
	const method = "notes.validateMany"
	var resp Items[Validation]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Validation]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CreateNote(ctx context.Context, params CreateNoteParams) (NoteCreated, error) {
	const method = "notes.create"
	var resp NoteCreated
	if err := c.call(ctx, method, params, &resp); err != nil {
		return NoteCreated{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CreateNotes(ctx context.Context, params CreateNotesParams) (Items[NoteCreated], error) {
	const method = "notes.createMany"
	var resp Items[NoteCreated]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[NoteCreated]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateNote(ctx context.Context, params UpdateNoteParams) (Note, error) {
	const method = "notes.update"
	var resp Note
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Note{}, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteNotes(ctx context.Context, params EntitiesParams) (Items[Ack], error) {
	const method = "notes.delete"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) PlanConversion(ctx context.Context, params PlanConversionParams) (ConversionPlan, error) {
	const method = "notes.planConversion"
	var resp ConversionPlan
	if err := c.call(ctx, method, params, &resp); err != nil {
		return ConversionPlan{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ConvertNotes(ctx context.Context, params ConvertNotesParams) (Count, error) {
	const method = "notes.convert"
	var resp Count
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Count{}, err
	}
	return resp, nil
}

func (c *AnkiClient) FindEmptyNotes(ctx context.Context, params FindEmptyNotesParams) (IdPage, error) {
	const method = "notes.findEmpty"
	var resp IdPage
	if err := c.call(ctx, method, params, &resp); err != nil {
		return IdPage{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListTags(ctx context.Context, params SessionParams) ([]string, error) {
	const method = "tags.list"
	var resp []string
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) AddTags(ctx context.Context, params TagsParams) (Items[Ack], error) {
	const method = "tags.add"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) RemoveTags(ctx context.Context, params TagsParams) (Items[Ack], error) {
	const method = "tags.remove"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ReplaceTags(ctx context.Context, params ReplaceTagsParams) (Count, error) {
	const method = "tags.replace"
	var resp Count
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Count{}, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteUnusedTags(ctx context.Context, params SessionParams) (Count, error) {
	const method = "tags.deleteUnused"
	var resp Count
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Count{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetCard(ctx context.Context, params EntityParams) (Card, error) {
	const method = "cards.get"
	var resp Card
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Card{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetCards(ctx context.Context, params EntitiesParams) (Items[Card], error) {
	const method = "cards.getMany"
	var resp Items[Card]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Card]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SearchCards(ctx context.Context, params SearchParams) (IdPage, error) {
	const method = "cards.search"
	var resp IdPage
	if err := c.call(ctx, method, params, &resp); err != nil {
		return IdPage{}, err
	}
	return resp, nil
}

func (c *AnkiClient) RenderCard(ctx context.Context, params EntityParams) (RenderedCard, error) {
	const method = "cards.render"
	var resp RenderedCard
	if err := c.call(ctx, method, params, &resp); err != nil {
		return RenderedCard{}, err
	}
	return resp, nil
}

func (c *AnkiClient) MoveCards(ctx context.Context, params MoveCardsParams) (Items[Ack], error) {
	const method = "cards.move"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SetSuspended(ctx context.Context, params SetSuspendedParams) (Items[Ack], error) {
	const method = "cards.setSuspended"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SetFlags(ctx context.Context, params SetFlagsParams) (Items[Ack], error) {
	const method = "cards.setFlags"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetScheduling(ctx context.Context, params CardParams) (SchedulingState, error) {
	const method = "scheduling.get"
	var resp SchedulingState
	if err := c.call(ctx, method, params, &resp); err != nil {
		return SchedulingState{}, err
	}
	return resp, nil
}

func (c *AnkiClient) PredictScheduling(ctx context.Context, params CardParams) (Predictions, error) {
	const method = "scheduling.predict"
	var resp Predictions
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Predictions{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetDue(ctx context.Context, params DueEligibilityParams) (DueEligibility, error) {
	const method = "scheduling.getDue"
	var resp DueEligibility
	if err := c.call(ctx, method, params, &resp); err != nil {
		return DueEligibility{}, err
	}
	return resp, nil
}

func (c *AnkiClient) AnswerCard(ctx context.Context, params AnswerParams) (AnswerResult, error) {
	const method = "scheduling.answer"
	var resp AnswerResult
	if err := c.call(ctx, method, params, &resp); err != nil {
		return AnswerResult{}, err
	}
	return resp, nil
}

func (c *AnkiClient) AnswerCards(ctx context.Context, params AnswerManyParams) (Items[AnswerResult], error) {
	const method = "scheduling.answerMany"
	var resp Items[AnswerResult]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[AnswerResult]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ResetCards(ctx context.Context, params ResetCardsParams) (Items[Ack], error) {
	const method = "scheduling.reset"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SetDueDate(ctx context.Context, params SetDueDateParams) (Items[Ack], error) {
	const method = "scheduling.setDueDate"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListDecks(ctx context.Context, params SessionParams) ([]Deck, error) {
	const method = "decks.list"
	var resp []Deck
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) GetDeck(ctx context.Context, params DeckParams) (Deck, error) {
	const method = "decks.get"
	var resp Deck
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Deck{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CreateDeck(ctx context.Context, params CreateDeckParams) (Deck, error) {
	const method = "decks.create"
	var resp Deck
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Deck{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateDeck(ctx context.Context, params UpdateDeckParams) (Deck, error) {
	const method = "decks.update"
	var resp Deck
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Deck{}, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteDecks(ctx context.Context, params DeleteDecksParams) (Items[Ack], error) {
	const method = "decks.delete"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetDeckCounts(ctx context.Context, params DeckCountsParams) (DeckCounts, error) {
	const method = "decks.getCounts"
	var resp DeckCounts
	if err := c.call(ctx, method, params, &resp); err != nil {
		return DeckCounts{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListDeckPresets(ctx context.Context, params SessionParams) ([]EntitySummary, error) {
	const method = "deckPresets.list"
	var resp []EntitySummary
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) GetDeckPreset(ctx context.Context, params PresetParams) (DeckPreset, error) {
	const method = "deckPresets.get"
	var resp DeckPreset
	if err := c.call(ctx, method, params, &resp); err != nil {
		return DeckPreset{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CreateDeckPreset(ctx context.Context, params CreatePresetParams) (DeckPreset, error) {
	const method = "deckPresets.create"
	var resp DeckPreset
	if err := c.call(ctx, method, params, &resp); err != nil {
		return DeckPreset{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateDeckPreset(ctx context.Context, params UpdatePresetParams) (OperationHandle, error) {
	const method = "deckPresets.update"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) AssignDeckPreset(ctx context.Context, params AssignPresetParams) (OperationHandle, error) {
	const method = "deckPresets.assign"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteDeckPreset(ctx context.Context, params DeletePresetParams) (OperationHandle, error) {
	const method = "deckPresets.delete"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListNoteTypes(ctx context.Context, params SessionParams) ([]EntitySummary, error) {
	const method = "noteTypes.list"
	var resp []EntitySummary
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) GetNoteType(ctx context.Context, params NoteTypeParams) (NoteType, error) {
	const method = "noteTypes.get"
	var resp NoteType
	if err := c.call(ctx, method, params, &resp); err != nil {
		return NoteType{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CreateNoteType(ctx context.Context, params CreateNoteTypeParams) (NoteType, error) {
	const method = "noteTypes.create"
	var resp NoteType
	if err := c.call(ctx, method, params, &resp); err != nil {
		return NoteType{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateNoteType(ctx context.Context, params UpdateNoteTypeParams) (NoteType, error) {
	const method = "noteTypes.update"
	var resp NoteType
	if err := c.call(ctx, method, params, &resp); err != nil {
		return NoteType{}, err
	}
	return resp, nil
}

func (c *AnkiClient) EditNoteTypeStructure(ctx context.Context, params EditStructureParams) (NoteType, error) {
	const method = "noteTypes.editStructure"
	var resp NoteType
	if err := c.call(ctx, method, params, &resp); err != nil {
		return NoteType{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetTemplateUsage(ctx context.Context, params NoteTypeParams) ([]TemplateUsage, error) {
	const method = "noteTypes.getTemplateUsage"
	var resp []TemplateUsage
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) ReplaceNoteTypeText(ctx context.Context, params ReplaceTextParams) (Count, error) {
	const method = "noteTypes.replaceText"
	var resp Count
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Count{}, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteUnusedNoteTypes(ctx context.Context, params DeleteNoteTypesParams) (Items[Ack], error) {
	const method = "noteTypes.deleteUnused"
	var resp Items[Ack]
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Items[Ack]{}, err
	}
	return resp, nil
}

func (c *AnkiClient) StoreMedia(ctx context.Context, params StoreMediaParams) (OperationHandle, error) {
	const method = "media.store"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetMedia(ctx context.Context, params MediaParams) (MediaFile, error) {
	const method = "media.get"
	var resp MediaFile
	if err := c.call(ctx, method, params, &resp); err != nil {
		return MediaFile{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListMedia(ctx context.Context, params ListMediaParams) ([]string, error) {
	const method = "media.list"
	var resp []string
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) DeleteMedia(ctx context.Context, params MediaParams) (Ack, error) {
	const method = "media.delete"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetMediaDirectory(ctx context.Context, params SessionParams) (MediaDirectory, error) {
	const method = "media.getDirectory"
	var resp MediaDirectory
	if err := c.call(ctx, method, params, &resp); err != nil {
		return MediaDirectory{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ListReviews(ctx context.Context, params ListReviewsParams) (ReviewPage, error) {
	const method = "reviews.list"
	var resp ReviewPage
	if err := c.call(ctx, method, params, &resp); err != nil {
		return ReviewPage{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetLatestReviewId(ctx context.Context, params LatestReviewParams) (LatestReviewId, error) {
	const method = "reviews.getLatestId"
	var resp LatestReviewId
	if err := c.call(ctx, method, params, &resp); err != nil {
		return LatestReviewId{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetTodayStatistics(ctx context.Context, params SessionParams) (ReviewCount, error) {
	const method = "statistics.getToday"
	var resp ReviewCount
	if err := c.call(ctx, method, params, &resp); err != nil {
		return ReviewCount{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetReviewDays(ctx context.Context, params SessionParams) ([]ReviewDay, error) {
	const method = "statistics.getReviewDays"
	var resp []ReviewDay
	if err := c.call(ctx, method, params, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AnkiClient) GetStatisticsSummary(ctx context.Context, params StatisticsParams) (StatisticsSummary, error) {
	const method = "statistics.getSummary"
	var resp StatisticsSummary
	if err := c.call(ctx, method, params, &resp); err != nil {
		return StatisticsSummary{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetStatisticsHtml(ctx context.Context, params StatisticsParams) (StatisticsHtml, error) {
	const method = "statistics.getHtml"
	var resp StatisticsHtml
	if err := c.call(ctx, method, params, &resp); err != nil {
		return StatisticsHtml{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ImportPackage(ctx context.Context, params ImportPackageParams) (OperationHandle, error) {
	const method = "packages.import"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ExportPackage(ctx context.Context, params ExportPackageParams) (OperationHandle, error) {
	const method = "packages.export"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) OpenBrowser(ctx context.Context, params OpenBrowserParams) (OperationHandle, error) {
	const method = "browser.open"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetBrowserSelection(ctx context.Context, params SessionParams) (BrowserState, error) {
	const method = "browser.getSelection"
	var resp BrowserState
	if err := c.call(ctx, method, params, &resp); err != nil {
		return BrowserState{}, err
	}
	return resp, nil
}

func (c *AnkiClient) SelectBrowser(ctx context.Context, params SelectBrowserParams) (OperationHandle, error) {
	const method = "browser.select"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) OpenNoteEditor(ctx context.Context, params OpenNoteEditorParams) (OperationHandle, error) {
	const method = "editors.openNote"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) OpenAddEditor(ctx context.Context, params OpenAddEditorParams) (OperationHandle, error) {
	const method = "editors.openAdd"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateDraft(ctx context.Context, params UpdateDraftParams) (OperationHandle, error) {
	const method = "editors.updateDraft"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetReviewer(ctx context.Context, params SessionParams) (ReviewerState, error) {
	const method = "reviewer.get"
	var resp ReviewerState
	if err := c.call(ctx, method, params, &resp); err != nil {
		return ReviewerState{}, err
	}
	return resp, nil
}

func (c *AnkiClient) StartReviewerTimer(ctx context.Context, params ReviewerTokenParams) (Ack, error) {
	const method = "reviewer.startTimer"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ShowReviewerQuestion(ctx context.Context, params ReviewerTokenParams) (Ack, error) {
	const method = "reviewer.showQuestion"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ShowReviewerAnswer(ctx context.Context, params ReviewerTokenParams) (Ack, error) {
	const method = "reviewer.showAnswer"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) AnswerReviewer(ctx context.Context, params ReviewerAnswerParams) (OperationHandle, error) {
	const method = "reviewer.answer"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) PlayReviewerAudio(ctx context.Context, params ReviewerTokenParams) (Ack, error) {
	const method = "reviewer.playAudio"
	var resp Ack
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Ack{}, err
	}
	return resp, nil
}

func (c *AnkiClient) Navigate(ctx context.Context, params NavigateParams) (OperationHandle, error) {
	const method = "application.navigate"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ImportFile(ctx context.Context, params ImportFileParams) (OperationHandle, error) {
	const method = "application.importFile"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) Exit(ctx context.Context) (ExitResponse, error) {
	const method = "application.exit"
	var resp ExitResponse
	if err := c.call(ctx, method, emptyPrams, &resp); err != nil {
		return ExitResponse{}, err
	}
	return resp, nil
}

func (c *AnkiClient) UpdateCard(ctx context.Context, params UpdateCardParams) (Card, error) {
	const method = "advanced.cards.update"
	var resp Card
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Card{}, err
	}
	return resp, nil
}

func (c *AnkiClient) ImportReviews(ctx context.Context, params ImportReviewsParams) (OperationHandle, error) {
	const method = "advanced.reviews.import"
	var resp OperationHandle
	if err := c.call(ctx, method, params, &resp); err != nil {
		return OperationHandle{}, err
	}
	return resp, nil
}

func (c *AnkiClient) GetOperation(ctx context.Context, params OperationParams) (Operation, error) {
	const method = "operations.get"
	var resp Operation
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Operation{}, err
	}
	return resp, nil
}

func (c *AnkiClient) CancelOperation(ctx context.Context, params OperationParams) (Cancellation, error) {
	const method = "operations.cancel"
	var resp Cancellation
	if err := c.call(ctx, method, params, &resp); err != nil {
		return Cancellation{}, err
	}
	return resp, nil
}
