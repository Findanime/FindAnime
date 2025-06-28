from openai import OpenAI
import json
client = OpenAI(
  api_key="key"
)

def AskAI(anime : str) -> dict :
  completion = client.chat.completions.create(
  model="gpt-4o-mini",
  store=True,
  messages=[
    {"role": "user", "content": f"give some animes like {anime} with their rating and correct banner Image and short description in a json format. No boiler plate texts give precisely only and only the json. No junk texts and not in codeblocks"}
  ]
)
  return json.loads(completion.choices[0].message.content)
