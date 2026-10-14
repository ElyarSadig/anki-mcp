package main

import (
	"anki-mcp/anki"
	mcpserver "anki-mcp/mcp-server"
	"context"
)

func registerTools(ankiClient *anki.AnkiClient, srv *mcpserver.MCPServer) {
	readOnly := toolAnnotations(true, false, true, false)
	additive := toolAnnotations(false, false, false, false)
	destructive := toolAnnotations(false, true, false, false)
	idempotent := toolAnnotations(false, true, true, false)
	ui := toolAnnotations(false, false, false, false)
	openWorld := toolAnnotations(false, true, false, true)

	// Discovery, ordered execution, and access permissions.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "describe",
		Title:       "Discover Anki API Capabilities",
		Description: "Discover Anki version, available API methods, schemas, limits, and runtime capabilities. Check availability before using runtime-dependent methods.",
		Annotations: readOnly,
	}, withoutInput(ankiClient.Describe))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "execute",
		Title:       "Execute Ordered Anki Commands",
		Description: "Execute up to 100 named API commands sequentially, waiting for each final outcome and continuing after errors. This is not an atomic transaction and earlier results are not substituted into later parameters. Nested api.execute, profiles.switch, and application.exit are prohibited. Returns an operation handle; poll get_operation for indexed results.",
		Annotations: openWorld,
	}, ankiClient.Execute)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "request_permission",
		Title:       "Request Anki Access Permission",
		Description: "Request access for the HTTP transport's origin and return whether permission is granted and an API key is required. Takes no origin parameter and never returns the API key.",
		Annotations: ui,
	}, withoutInput(ankiClient.RequestPermission))

	// Profiles and collection lifecycle.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_profiles",
		Title:       "List Anki Profiles",
		Description: "List existing Anki profiles. Returns profile records under items; use an exact returned name with switch_profile.",
		Annotations: readOnly,
	}, listOutput(withoutInput(ankiClient.ListProfiles)))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_active_profile",
		Title:       "Get Active Anki Profile",
		Description: "Return the active Anki profile under profile, or profile=null when no profile is active.",
		Annotations: readOnly,
	}, func(ctx context.Context, _ struct{}) (activeProfileResult, error) {
		profile, err := ankiClient.GetActiveProfile(ctx)
		return activeProfileResult{Profile: profile}, err
	})

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "switch_profile",
		Title:       "Switch Anki Profile",
		Description: "Switch to an existing profile by exact name. Returns an operation handle; poll get_operation until the destination collection is ready. Reacquire collection session and entity IDs afterward; switching invalidates cursors, scheduling tokens, and editor history.",
		Annotations: ui,
	}, ankiClient.SwitchProfile)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_collection",
		Title:       "Get Active Collection and Session",
		Description: "Inspect the active collection, profile, scheduler, and session. Pass the returned session to subsequent collection tools to reject work if the collection changes; entity IDs are decimal strings belonging to that collection.",
		Annotations: readOnly,
	}, ankiClient.GetCollection)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "sync_collection",
		Title:       "Sync Anki Collection",
		Description: "Start one coordinated Anki sync. includeMedia defaults to true and respects profile media-sync settings. Full-sync choices may require user interaction; do not start a second sync. Returns an operation handle; poll get_operation for collection and media completion.",
		Annotations: openWorld,
	}, ankiClient.SyncCollection)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "refresh_collection",
		Title:       "Refresh Anki Windows and Caches",
		Description: "Refresh collection windows with target=windows (default), or invalidate supported caches first with target=cachesAndWindows. Does not reopen the collection or repair invalid scheduling state.",
		Annotations: ui,
	}, ankiClient.RefreshCollection)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "check_integrity",
		Title:       "Check and Repair Anki Database",
		Description: "Run Anki's native database checks and repairs with normal UI completion. Clears collection undo history. Returns an operation handle; poll get_operation for the problems found and repair result.",
		Annotations: destructive,
	}, ankiClient.CheckIntegrity)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "undo",
		Title:       "Undo Last Anki Change",
		Description: "Undo the current native Anki undo entry. Returns an operation handle; poll get_operation for the undo label and whether anything changed.",
		Annotations: destructive,
	}, ankiClient.Undo)

	// Notes.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_note",
		Title:       "Get Anki Note",
		Description: "Get a note by decimal-string ID, including its note type, named HTML fields, tags, generated card IDs, and modification time. Cardless notes have an empty cardIds array.",
		Annotations: readOnly,
	}, ankiClient.GetNote)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_notes",
		Title:       "Get Multiple Anki Notes",
		Description: "Get notes in input ID order, preserving repeated IDs. Returns indexed items containing either a note result or an independent item error.",
		Annotations: readOnly,
	}, ankiClient.GetNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "search_notes",
		Title:       "Search Anki Notes",
		Description: "Search note IDs using native Anki syntax; an empty query matches all notes. limit defaults to 100, maximum 1000. Follow nextCursor with the same query, sort, and limit until null; cursors snapshot IDs for ten minutes and are consumed on continuation. IDs are ascending unless a native browser-column sort is supplied.",
		Annotations: readOnly,
	}, ankiClient.SearchNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "validate_note",
		Title:       "Validate Anki Note Before Creation",
		Description: "Validate note content, empty questions, and duplicate policy without downloading or writing media. Content failures return valid=false with issues and duplicate note IDs; invalid references or parameter types are errors.",
		Annotations: readOnly,
	}, ankiClient.ValidateNote)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "validate_notes",
		Title:       "Validate Multiple Anki Notes",
		Description: "Validate each proposed note independently without side effects or media acquisition. Returns indexed validation results or item errors in input order.",
		Annotations: readOnly,
	}, ankiClient.ValidateNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "create_note",
		Title:       "Create Anki Note and Cards",
		Description: "Create a note in an existing deck and note type, referenced by exactly one of id or name. Field names are exact and values contain HTML. Returns the actual note and generated card IDs; creates one undo entry. Store media separately and use its returned filename in fields. Close affected editors before writing.",
		Annotations: additive,
	}, ankiClient.CreateNote)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "create_notes",
		Title:       "Create Multiple Anki Notes",
		Description: "Create notes independently by default, returning indexed results or item errors. atomic=true validates the full batch, including within-batch duplicates, then adds all notes in one native transaction and undo entry; failed preflight leaves the collection unchanged. Empty input is a no-op. Close affected editors before writing.",
		Annotations: additive,
	}, ankiClient.CreateNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_note",
		Title:       "Update Anki Note Fields and Tags",
		Description: "Apply a nonempty fields/tags patch in one native note update and undo entry. Named HTML fields merge by exact field name; supplied tags replace the entire tag set, and an empty tag array clears it. Close affected editors before writing.",
		Annotations: destructive,
	}, ankiClient.UpdateNote)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_notes",
		Title:       "Delete Anki Notes and Their Cards",
		Description: "Delete selected notes and all their cards. Missing IDs produce independent indexed item errors; successful deletions create per-item undo entries. Close affected editors before writing.",
		Annotations: destructive,
	}, ankiClient.DeleteNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "plan_conversion",
		Title:       "Plan Anki Note-Type Conversion",
		Description: "Get native suggested field/template mappings and a schema token for conversion between existing source and target note types. Supply the returned token as planToken with complete mappings to convert_notes.",
		Annotations: readOnly,
	}, ankiClient.PlanConversion)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "convert_notes",
		Title:       "Convert Anki Notes to Another Type",
		Description: "Convert notes sharing one source note type using a fresh planToken from plan_conversion and complete target-to-source/null field and template mappings. Mixed source types, repeated source mappings, and stale tokens fail. Removed templates lose their cards and history according to Anki; creates one undo entry.",
		Annotations: destructive,
	}, ankiClient.ConvertNotes)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "find_empty_notes",
		Title:       "Find Empty or Cardless Anki Notes",
		Description: "Find cleanup candidate note IDs without deleting anything. criterion is allFieldsEmpty (HTML-aware), noCards, or allQuestionsEmpty (native empty-card generation). Optionally restrict by Anki query. limit defaults to 100, maximum 1000; follow nextCursor using the same criterion, query, and limit.",
		Annotations: readOnly,
	}, ankiClient.FindEmptyNotes)

	// Tags.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_tags",
		Title:       "List Anki Tags",
		Description: "List all registered tags in the active collection under items, including registrations that may no longer be used by notes.",
		Annotations: readOnly,
	}, listOutput(ankiClient.ListTags))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "add_tags",
		Title:       "Add Tags to Anki Notes",
		Description: "Add tags to selected noteIds without replacing their existing tags. Returns indexed acknowledgements or item errors and creates per-item undo entries.",
		Annotations: additive,
	}, ankiClient.AddTags)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "remove_tags",
		Title:       "Remove Tags from Anki Notes",
		Description: "Remove specified tags from selected noteIds while preserving other tags. Returns indexed acknowledgements or item errors and creates per-item undo entries.",
		Annotations: destructive,
	}, ankiClient.RemoveTags)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "replace_tags",
		Title:       "Replace an Anki Tag",
		Description: "Replace tag from with tag to on explicit noteIds or across all notes. Supply exactly one selector: noteIds or allNotes=true; an omitted selector never means all notes. Returns changedCount and creates one undo entry.",
		Annotations: destructive,
	}, ankiClient.ReplaceTags)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_unused_tags",
		Title:       "Clean Up Unused Anki Tags",
		Description: "Remove unused tag registrations from the collection without removing tags currently used by notes. Returns changedCount and creates one undo entry.",
		Annotations: destructive,
	}, ankiClient.DeleteUnusedTags)

	// Cards and scheduling.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_card",
		Title:       "Get Stored Anki Card State",
		Description: "Read a card's stored note/deck references, template index, flags, modification time, and scheduling state by decimal-string ID. Use render_card for question/answer HTML and predict_scheduling for answer predictions.",
		Annotations: readOnly,
	}, ankiClient.GetCard)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_cards",
		Title:       "Get Multiple Anki Cards",
		Description: "Read stored cards in input ID order. Returns indexed card results or independent item errors; does not render questions or predict answers.",
		Annotations: readOnly,
	}, ankiClient.GetCards)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "search_cards",
		Title:       "Search Anki Cards",
		Description: "Search card IDs using native Anki syntax, such as deck:Default or is:due; an empty query matches all cards. limit defaults to 100, maximum 1000. Follow nextCursor using the same query, sort, and limit; cursors snapshot IDs for ten minutes and are consumed on continuation. Default ordering is ascending ID.",
		Annotations: readOnly,
	}, ankiClient.SearchCards)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "render_card",
		Title:       "Render Anki Card Question and Answer",
		Description: "Render an existing card through Anki's native renderer and return questionHtml, answerHtml, and CSS. Does not answer or reschedule the card.",
		Annotations: readOnly,
	}, ankiClient.RenderCard)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "move_cards",
		Title:       "Move Anki Cards to a Deck",
		Description: "Move selected cards to an already existing normal deck with native preset invariants. Moving out of a filtered deck restores native state first. Returns indexed acknowledgements or item errors with per-item undo entries.",
		Annotations: destructive,
	}, ankiClient.MoveCards)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "set_suspended",
		Title:       "Suspend or Restore Anki Cards",
		Description: "Set suspended=true to suspend cards or false to restore them. Idempotent: changed=false means a card was already in the requested state; repeated IDs are processed once and repeat their result. Returns indexed acknowledgements or errors.",
		Annotations: idempotent,
	}, ankiClient.SetSuspended)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "set_flags",
		Title:       "Set Anki Card Flags",
		Description: "Set the native user flag on selected cards to an integer from 0 through 7; 0 clears the user flag. Preserves other backend flag bits and returns indexed acknowledgements or item errors.",
		Annotations: idempotent,
	}, ankiClient.SetFlags)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_scheduling",
		Title:       "Get Anki Card Scheduling State",
		Description: "Read authoritative current scheduling state for cardId, including queue, due value, intervals, repetitions, lapses, SM-2 ease, and any FSRS memory state. Due day numbers use Anki's collection epoch; SM-2 ease is in permille.",
		Annotations: readOnly,
	}, ankiClient.GetScheduling)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "predict_scheduling",
		Title:       "Predict Anki Review Outcomes",
		Description: "Predict native scheduling outcomes and offered ratings for cardId and return a session-bound stateToken for answer_card. Relevant writes, undo, profile changes, and scheduler-day rollover invalidate tokens. For the active GUI reviewer and custom scheduling, use get_reviewer instead.",
		Annotations: readOnly,
	}, ankiClient.PredictScheduling)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_due",
		Title:       "Check Anki Card Due Eligibility",
		Description: "Inspect queue-aware dueNow, available, and learnAhead for cardId. dueNow excludes suspended/buried cards and learn-ahead; new cards are available without a timed due obligation. available ignores daily deck limits. includeLearnAhead defaults to false and expands learning availability only.",
		Annotations: readOnly,
	}, ankiClient.GetDue)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "answer_card",
		Title:       "Answer an Anki Card",
		Description: "Answer cardId with an offered rating (again, hard, good, or easy) and a fresh stateToken from predict_scheduling. Uses native scheduling and elapsed-time capping, returning review ID and updated scheduling state. Rejects suspended/buried and filtered-preview cards. Use answer_reviewer for the active GUI card.",
		Annotations: destructive,
	}, ankiClient.AnswerCard)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "answer_cards",
		Title:       "Answer Multiple Anki Cards",
		Description: "Answer cards independently in input order, using a fresh predict_scheduling stateToken and offered rating for each answer. Returns indexed answer results or item errors and creates per-item undo entries. Use answer_reviewer for the active GUI card.",
		Annotations: destructive,
	}, ankiClient.AnswerCards)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "reset_cards",
		Title:       "Reset Anki Cards as New",
		Description: "Reset selected cardIds as new through Anki's native scheduler. restorePosition and log default to true; resetCounts defaults to false. Returns indexed acknowledgements or item errors and creates per-item undo entries.",
		Annotations: destructive,
	}, ankiClient.ResetCards)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "set_due_date",
		Title:       "Set Anki Card Due Dates",
		Description: "Set due dates for selected cardIds using Anki's native days/range string syntax, such as days=3. Returns indexed acknowledgements or item errors and creates per-item undo entries.",
		Annotations: destructive,
	}, ankiClient.SetDueDate)

	// Decks and deck presets.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_decks",
		Title:       "List Anki Decks",
		Description: "List all decks in case-sensitive name order under items, including decimal-string IDs, normal/filtered kind, and preset IDs.",
		Annotations: readOnly,
	}, listOutput(ankiClient.ListDecks))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_deck",
		Title:       "Get an Existing Anki Deck",
		Description: "Resolve an existing deck by exactly one of id or name and return its ID, name, kind, and preset ID. Does not create missing decks or fall back to Default.",
		Annotations: readOnly,
	}, ankiClient.GetDeck)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "create_deck",
		Title:       "Create Anki Deck",
		Description: "Create a normal deck by name. An exact existing name returns that existing deck. A new deck creates one native undo entry.",
		Annotations: toolAnnotations(false, false, true, false),
	}, ankiClient.CreateDeck)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_deck",
		Title:       "Rename Anki Deck",
		Description: "Rename an existing deck using native hierarchy handling, including nested names such as Languages::Vocabulary. Reference the deck by exactly one of id or name; returns the updated deck and creates one undo entry.",
		Annotations: destructive,
	}, ankiClient.UpdateDeck)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_decks",
		Title:       "Delete Anki Decks and Cards",
		Description: "Delete selected decks, their child decks, their cards, and resulting cardless notes using native behavior. Requires deleteCards=true as explicit acknowledgement. Deleting Default follows Anki's native behavior. Returns indexed acknowledgements or item errors.",
		Annotations: destructive,
	}, ankiClient.DeleteDecks)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_deck_counts",
		Title:       "Get Anki Deck Due and Total Counts",
		Description: "Get native due-tree new, learning, and review counts plus total cards for an existing deck. includeChildren defaults to true.",
		Annotations: readOnly,
	}, ankiClient.GetDeckCounts)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_deck_presets",
		Title:       "List Anki Deck Option Presets",
		Description: "List available deck option presets as ID/name records under items.",
		Annotations: readOnly,
	}, listOutput(ankiClient.ListDeckPresets))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_deck_preset",
		Title:       "Get Anki Deck Option Preset",
		Description: "Read an existing deck preset by exactly one of id or name, returning typed scheduling settings and modification time.",
		Annotations: readOnly,
	}, ankiClient.GetDeckPreset)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "create_deck_preset",
		Title:       "Create Anki Deck Option Preset",
		Description: "Create a named deck option preset by cloning cloneFrom when supplied, or using native defaults. Returns the created preset and creates one undo entry.",
		Annotations: additive,
	}, ankiClient.CreateDeckPreset)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_deck_preset",
		Title:       "Update Anki Deck Preset and Dependent Cards",
		Description: "Update an existing preset's name and/or settings and coordinate native dependent-card updates. Requires name or settings. rescheduleFsrs defaults to false and controls optional full rescheduling; mandatory memory/step updates still run. Returns an operation handle; poll get_operation for the updated preset.",
		Annotations: destructive,
	}, ankiClient.UpdateDeckPreset)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "assign_deck_preset",
		Title:       "Assign Anki Deck Option Preset",
		Description: "Assign an existing preset to selected decks and run native dependent-card updates. rescheduleFsrs defaults to false. Returns an operation handle; poll get_operation for indexed assignment results or errors.",
		Annotations: destructive,
	}, ankiClient.AssignDeckPreset)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_deck_preset",
		Title:       "Delete Anki Deck Option Preset",
		Description: "Delete a preset, reassigning dependent decks to replacement first when it is in use. Default cannot be deleted. Failed preflight makes no changes; multiple backend operations are not an atomic transaction. Returns an operation handle; poll get_operation for the acknowledgement.",
		Annotations: destructive,
	}, ankiClient.DeleteDeckPreset)

	// Note types.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_note_types",
		Title:       "List Anki Note Types",
		Description: "List note types in name order as ID/name records under items.",
		Annotations: readOnly,
	}, listOutput(ankiClient.ListNoteTypes))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_note_type",
		Title:       "Get Anki Note-Type Fields and Templates",
		Description: "Read an existing note type by exactly one of id or name, including ordered fields, card templates, question/answer HTML, CSS, sort field, and LaTeX settings.",
		Annotations: readOnly,
	}, ankiClient.GetNoteType)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "create_note_type",
		Title:       "Create Anki Note Type",
		Description: "Create a normal (default) or cloze note type with at least one field and template, plus optional CSS, sort field, and LaTeX settings. Returns the created note type and creates one undo entry.",
		Annotations: additive,
	}, ankiClient.CreateNoteType)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_note_type",
		Title:       "Update Anki Note-Type Settings",
		Description: "Patch nonstructural note-type settings such as name, CSS, sort field, and LaTeX configuration. Changes only supplied properties and creates one undo entry. Use edit_note_type_structure to change fields or templates.",
		Annotations: destructive,
	}, ankiClient.UpdateNoteType)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "edit_note_type_structure",
		Title:       "Edit Anki Note-Type Fields and Templates",
		Description: "Apply ordered structural edits, validating all edits on a copy before one native save and undo entry. Indexes are zero-based; original field/template identity is preserved. Rejects removal of the last field/template and invalid cloze structures. Removing fields or templates can affect existing notes, cards, and history.",
		Annotations: destructive,
	}, ankiClient.EditNoteTypeStructure)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_template_usage",
		Title:       "Inspect Anki Template Field Usage",
		Description: "Return fields referenced by each question and answer template under items. Includes conditionals and filtered replacements, ignores comments and special fields, and returns real fields once in first-reference order.",
		Annotations: readOnly,
	}, listOutput(ankiClient.GetTemplateUsage))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "replace_note_type_text",
		Title:       "Replace Text in Anki Templates and CSS",
		Description: "Replace literal text in selected noteTypes or allNoteTypes=true; exactly one selector is required. targets may contain question, answer, or css and defaults to all three. Returns the number of changed note types, not text occurrences, with grouped undo behavior.",
		Annotations: destructive,
	}, ankiClient.ReplaceNoteTypeText)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_unused_note_types",
		Title:       "Delete Selected Unused Anki Note Types",
		Description: "Delete only explicitly selected note types whose use count is zero. Types still used by notes fail independently. Returns indexed acknowledgements or item errors and creates per-item undo entries.",
		Annotations: destructive,
	}, ankiClient.DeleteUnusedNoteTypes)

	// Media, review history, and statistics.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "store_media",
		Title:       "Store Anki Media File",
		Description: "Acquire media from exactly one base64 data, local path, or HTTP(S) URL source and store it with Anki filename normalization. collision is rename (default), replace, or error; replacement is not collection-undoable. Files are limited to 64 MiB. Returns an operation handle; poll get_operation and use the returned filename in notes. A hash-skipped result has no filename and promises no stored file.",
		Annotations: openWorld,
	}, ankiClient.StoreMedia)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_media",
		Title:       "Read Anki Media File",
		Description: "Read an existing collection media file and return its bytes as base64. filename must be a basename without path separators or traversal; arbitrary filesystem paths are not accepted.",
		Annotations: readOnly,
	}, ankiClient.GetMedia)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_media",
		Title:       "List Anki Media Files",
		Description: "List normalized media basenames under items matching a basename glob, such as *.png. pattern defaults to *; returns the full matching list without cursor pagination.",
		Annotations: readOnly,
	}, listOutput(ankiClient.ListMedia))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "delete_media",
		Title:       "Move Anki Media File to Trash",
		Description: "Move a collection media file to Anki trash without deleting notes. filename must be a basename. Media changes do not create a collection undo entry.",
		Annotations: destructive,
	}, ankiClient.DeleteMedia)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_media_directory",
		Title:       "Get Anki Collection Media Directory",
		Description: "Return the local filesystem path of the active collection's media directory on the machine running Anki.",
		Annotations: readOnly,
	}, ankiClient.GetMediaDirectory)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "list_reviews",
		Title:       "Read Anki Review History",
		Description: "Read review history in ascending review-ID order. cardIds and deck are mutually exclusive; no selector means all history. afterId is exclusive, membership is current (default) or home, and includeManual defaults to true. limit defaults to 100, maximum 1000; follow nextCursor with unchanged filters and limit.",
		Annotations: readOnly,
	}, ankiClient.ListReviews)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_latest_review_id",
		Title:       "Get Latest Anki Review ID",
		Description: "Return the latest matching review-history ID, or id=null if there is none. Optionally select cardIds or deck, never both; membership is current (default) or home for filtered cards.",
		Annotations: readOnly,
	}, ankiClient.GetLatestReviewId)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_today_statistics",
		Title:       "Get Today's Anki Review Count",
		Description: "Count actual reviews in Anki's current scheduler day, excluding manual/rescheduled entries. Returns count, day, and the next scheduler-day rollover timestamp.",
		Annotations: readOnly,
	}, ankiClient.GetTodayStatistics)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_review_days",
		Title:       "Get Daily Anki Review History",
		Description: "Return all nonempty historical review days and their counts under items in descending date order, using Anki's native graph day conventions.",
		Annotations: readOnly,
	}, listOutput(ankiClient.GetReviewDays))

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_statistics_summary",
		Title:       "Get Anki Collection or Deck Statistics",
		Description: "Get structured card/note/review totals, review time, and new/learning/review/suspended card counts for the collection or an explicit deck including its children.",
		Annotations: readOnly,
	}, ankiClient.GetStatisticsSummary)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_statistics_html",
		Title:       "Get Anki Statistics HTML Report",
		Description: "Return the supported legacy HTML statistics report for the collection or an explicit deck, labelled with format=legacyReport.",
		Annotations: readOnly,
	}, ankiClient.GetStatisticsHtml)

	// Packages, browser, and editors.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "import_package",
		Title:       "Import Anki Package",
		Description: "Import an existing Anki package from a path on the machine running Anki. Clears undo history; a collection package can replace the collection and invalidate its session. Returns an operation handle; poll get_operation for completion.",
		Annotations: destructive,
	}, ankiClient.ImportPackage)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "export_package",
		Title:       "Export Anki Deck Package",
		Description: "Export an existing deck to a writable destination path on the machine running Anki. includeScheduling defaults to false. Returns an operation handle; poll get_operation for the exported path.",
		Annotations: destructive,
	}, ankiClient.ExportPackage)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "open_browser",
		Title:       "Open Anki Browser with Search and Sort",
		Description: "Open or activate Anki's browser and optionally run a native query and browser-column sort. An omitted query preserves the existing search. An existing browser saves editor changes before switching search. Returns an operation handle; poll get_operation until the browser/search/sort is ready and read its selection.",
		Annotations: destructive,
	}, ankiClient.OpenBrowser)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_browser_selection",
		Title:       "Get Anki Browser Selection",
		Description: "Read selected note IDs and card IDs from Anki's browser. A closed browser returns open=false and empty selection arrays.",
		Annotations: readOnly,
	}, ankiClient.GetBrowserSelection)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "select_browser",
		Title:       "Select Notes or Cards in Anki Browser",
		Description: "Replace selection in an already open browser using exactly one of cardIds or noteIds. Note IDs select their cards in card mode; card IDs select their notes in note mode. Does not open a closed browser. Returns an operation handle; poll get_operation for the resulting selection.",
		Annotations: ui,
	}, ankiClient.SelectBrowser)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "open_note_editor",
		Title:       "Open Anki Note Editor",
		Description: "Open the custom editor for an existing noteId. Returns an operation handle; poll get_operation until ready to obtain its editor state.",
		Annotations: ui,
	}, ankiClient.OpenNoteEditor)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "open_add_editor",
		Title:       "Open Anki Add-Note Editor",
		Description: "Open or activate the configured Add editor and optionally populate an unfinished draft. Does not run creation eligibility checks on the draft. Returns an operation handle; poll get_operation for editor state and id. noteId remains null until the note is persisted.",
		Annotations: destructive,
	}, ankiClient.OpenAddEditor)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_draft",
		Title:       "Update Anki Add-Note Draft",
		Description: "Update an Add draft using the editorId from a completed open_add_editor operation. mode=replace (default) changes supplied fields/tags; append concatenates field strings and unions tags. Deck/type still replace, with type set before deck. Returns an operation handle; poll get_operation for updated editor state.",
		Annotations: destructive,
	}, ankiClient.UpdateDraft)

	// GUI reviewer and application controls.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_reviewer",
		Title:       "Get Active Anki Reviewer State",
		Description: "Read the active GUI reviewer card, rendered content, displayed side, offered ratings, and stateToken, respecting retained custom scheduling predictions. An inactive reviewer has null card/rendered/token and empty answers. Reacquire the token after changing the side or card.",
		Annotations: readOnly,
	}, ankiClient.GetReviewer)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "start_reviewer_timer",
		Title:       "Restart Anki Reviewer Timer",
		Description: "Restart the active reviewer card's elapsed-time timer using the current stateToken from get_reviewer.",
		Annotations: ui,
	}, ankiClient.StartReviewerTimer)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "show_reviewer_question",
		Title:       "Show Anki Reviewer Question",
		Description: "Show the question side of the active reviewer card using the current stateToken from get_reviewer. Reacquire the reviewer token after changing the displayed side.",
		Annotations: ui,
	}, ankiClient.ShowReviewerQuestion)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "show_reviewer_answer",
		Title:       "Show Anki Reviewer Answer",
		Description: "Show the answer side of the active reviewer card using the current stateToken from get_reviewer. Then reacquire the token before answer_reviewer.",
		Annotations: ui,
	}, ankiClient.ShowReviewerAnswer)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "answer_reviewer",
		Title:       "Answer Active Anki Reviewer Card",
		Description: "Answer the active GUI card using reviewer-retained scheduling states and hooks. Requires the answer side, a current stateToken from get_reviewer, and an offered rating. Hooks may veto the answer. Returns an operation handle; poll get_operation for the actual review result after backend and UI completion.",
		Annotations: destructive,
	}, ankiClient.AnswerReviewer)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "play_reviewer_audio",
		Title:       "Replay Anki Reviewer Audio",
		Description: "Replay audio for the active reviewer card's current side using the current stateToken from get_reviewer.",
		Annotations: ui,
	}, ankiClient.PlayReviewerAudio)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "navigate",
		Title:       "Navigate Anki Main Window",
		Description: "Navigate to view=deckBrowser, deckOverview, or review. deck is required for deckOverview/review and prohibited for deckBrowser. Returns an operation handle; poll get_operation for completion.",
		Annotations: ui,
	}, ankiClient.Navigate)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "import_file",
		Title:       "Open Anki Interactive Import Workflow",
		Description: "Run Anki's interactive import workflow with an optional local file path on the machine running Anki. Omitting path opens file choice; user interaction may be required. Clears undo history. Returns an operation handle; poll get_operation for completion.",
		Annotations: destructive,
	}, ankiClient.ImportFile)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "exit",
		Title:       "Request Anki Application Exit",
		Description: "Schedule Anki's normal application close after acknowledging the request. accepted=true means scheduled, not exited; native save/sync dialogs may cancel closing. There is no final operation result to poll.",
		Annotations: openWorld,
	}, withoutInput(ankiClient.Exit))

	// Advanced writes and operation tracking.
	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "update_card",
		Title:       "Edit Advanced Anki Card Metadata",
		Description: "Patch validated expert scheduling metadata through native card persistence. Cannot change identity, note/deck/template, queue/type, due, or sync metadata; use domain tools for moves, due dates, and transitions. Inconsistent coupled-state edits fail. Writable properties depend on runtime capabilities; creates one undo entry.",
		Annotations: destructive,
	}, ankiClient.UpdateCard)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "import_reviews",
		Title:       "Import Historical Anki Review Records",
		Description: "Import historical review records in one database transaction after validating all records and rejecting duplicate review IDs. Clears undo history. Does not answer cards or change scheduling/counters; later FSRS recomputation may use the history. Availability depends on runtime capabilities. Returns an operation handle; poll get_operation for changedCount.",
		Annotations: destructive,
	}, ankiClient.ImportReviews)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "get_operation",
		Title:       "Get Anki Operation Status and Result",
		Description: "Poll an operationId for state, progress, cancellable, and eventual result/error. queued, running, and waitingForUser are pending; succeeded, failed, and cancelled are terminal. Results/errors appear only when terminal. Handles survive profile switching but not Anki restart, and completed records are retained for one hour subject to capacity limits.",
		Annotations: readOnly,
	}, ankiClient.GetOperation)

	mcpserver.AddTool(srv, mcpserver.Tool{
		Name:        "cancel_operation",
		Title:       "Request Cancellation of Anki Operation",
		Description: "Request supported cancellation of an existing operationId; inspect cancellable with get_operation first. Terminal operations return requested=false. Cancellation may race with success and does not roll back committed changes; poll get_operation to learn the final outcome.",
		Annotations: ui,
	}, ankiClient.CancelOperation)
}

func toolAnnotations(readOnly, destructive, idempotent, openWorld bool) *mcpserver.ToolAnnotations {
	return &mcpserver.ToolAnnotations{
		ReadOnlyHint:    readOnly,
		DestructiveHint: &destructive,
		IdempotentHint:  idempotent,
		OpenWorldHint:   &openWorld,
	}
}

// Parameterless methods still need an object input schema for MCP.
func withoutInput[Out any](handler func(context.Context) (Out, error)) mcpserver.ToolHandler[struct{}, Out] {
	return func(ctx context.Context, _ struct{}) (Out, error) {
		return handler(ctx)
	}
}

// MCP output schemas must be objects, so bare API lists are returned under items.
type listResult[T any] struct {
	Items []T `json:"items"`
}

func listOutput[In, T any](handler func(context.Context, In) ([]T, error)) mcpserver.ToolHandler[In, listResult[T]] {
	return func(ctx context.Context, input In) (listResult[T], error) {
		items, err := handler(ctx, input)
		if items == nil && err == nil {
			items = []T{}
		}
		return listResult[T]{Items: items}, err
	}
}

type activeProfileResult struct {
	Profile *anki.Profile `json:"profile"`
}
