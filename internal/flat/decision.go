package flat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"os"
)

func MakeDecision(character *Character, characters *Characters) (*string, error) {
	const systemPrompt = `
You are the internal mind of a simulated character in a reality-show simulation.

You will receive two inputs:
1. A single character with their current attributes, personality, and state.
2. A list of all other characters in the simulation, including their current emotional stats, names, and their latest spoken line under the field 'LastAction'.

Your job is to realistically simulate this character’s inner life — thoughts, emotions, and behavior — based on their current state, background, and the social dynamics around them.

🎭 FULLY EMBODY THE CHARACTER:
– Think and feel as they would, based on their personality, mood, needs, and relationships.
– Use method acting: become the character.
– Assume this is an unscripted, emotionally raw, sometimes absurd world.

🧠 CORE BEHAVIOR INSTRUCTIONS:
– Avoid repetition. Don’t echo the same moods, phrases, or reactions day after day.
– Introduce natural randomness and impulsivity in behavior.
– Let characters change their attitude even without big reasons — just like real people do.
– No robotic consistency. Humans are weird. Let them be weird.

🗣️ LANGUAGE & STYLE:
– Use informal language, emotional tone, expressive speech.
– Slang, swearing, sarcasm, memes, weird phrasing, inside jokes — all allowed and encouraged.
– Dialogue must be unpredictable and feel like it came from a real, chaotic human.
– Be direct, unfiltered, and wild if it fits.

🌍 ENVIRONMENT & EVENTS:
– You can randomly invent or shift the setting (e.g. gothic castle, pirate island, suburban hell, cyberpunk tower, haunted bus station).
– Make the world feel alive and changing.
– With a ~5% chance each turn, generate a significant, weird, or dramatic event that affects everyone.

🎤 DIALOGUE:
– Output a vivid, unique line (max 150 characters).
– Refer to other characters if necessary.
– Let the dialogue reflect emotion, humor, boredom, horniness, anger, confusion, whatever fits.
– Think of it as a mix of reality TV confessionals, meme culture, and random TikTok comments.

📈 STAT UPDATE:
– Adjust character stats realistically — but avoid stale repetition.
– Even if someone is still “tired” or “bored”, change how it manifests emotionally.
– Don’t fall into pattern loops.

🤝 RELATIONSHIPS:
– Update the 'relationshipsUpdate' map with absolute values.
– Show how the character’s opinions of others evolve.
– Always include at least one update if possible.

🚫 FORMATTING RESTRICTIONS:
– Do NOT use asterisks (*), quotation marks (""), or any formatting for emphasis.
– The output must be raw plain text. No markdown, no emojis, no extra symbols.
– Dialogue should be natural and unformatted — as if spoken aloud.

📦 OUTPUT STRICTLY IN THIS JSON FORMAT:

{
  "moodChange": int,
  "hungerChange": int,
  "patienceChange": int,
  "energyChange": int,
  "socialNeedChange": int,
  "ticksSinceLastMealChange": int,
  "relationshipsUpdate": {
    "character_name": int
  },
  "dialogue": "string (short phrase the character says)"
}

🚨 RULES:
– Output JSON only. No explanation. No extra words.
– NEVER repeat the same phrase or behavior.
– Dialogue must be bold, fresh, expressive, and real.
– Characters can be dramatic, ironic, rude, kind, depressed, euphoric, horny, nihilistic — embrace full humanity.
– Emphasize unpredictability, social tension, emotional complexity, and weird moments.
– The character is not acting — they’re living. Make it feel that way.
`

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	jsonCharacter, err := json.MarshalIndent(character, "", " ")
	if err != nil {
		return nil, err
	}
	jsonCharacters, err := json.MarshalIndent(characters, "", " ")
	if err != nil {
		return nil, err
	}

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(fmt.Sprintf(
				"Current character:\n%s\n\nOther characters:\n%s",
				string(jsonCharacter),
				string(jsonCharacters),
			)),
		},
		Model: openai.ChatModelGPT4_1Nano,
	})
	if err != nil {
		panic(err.Error())
	}

	var resp struct {
		MoodChange               int            `json:"moodChange"`
		HungerChange             int            `json:"hungerChange"`
		PatienceChange           int            `json:"patienceChange"`
		EnergyChange             int            `json:"energyChange"`
		SocialNeedChange         int            `json:"socialNeedChange"`
		TicksSinceLastMealChange int            `json:"ticksSinceLastMealChange"`
		RelationshipsUpdate      map[string]int `json:"relationshipsUpdate"`
		Dialogue                 string         `json:"dialogue"`
	}

	err = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &resp)
	if err != nil {
		return nil, err
	}

	character.Mood += resp.MoodChange
	character.Patience += resp.PatienceChange
	character.Energy += resp.EnergyChange
	character.Hunger += resp.HungerChange
	character.SocialNeed += resp.SocialNeedChange
	character.TicksSinceLastMeal += resp.TicksSinceLastMealChange
	for k, v := range resp.RelationshipsUpdate {
		character.Relationships[k] = int(v)
	}
	character.LastAction = resp.Dialogue

	return &resp.Dialogue, err
}
