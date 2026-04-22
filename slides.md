---
author: David Gageot & Djordje Lukić
paging: Slide %d / %d
theme: ./theme.json
---

# Docker Agent, le Couteau Suisse Agentique de Docker.

L’IA révolutionne nos métiers autour du développement.

Chez Docker, nous jouons les éclaireurs et changeons totalement
notre façon de travailler. 

Pour nous accompagner, nous avons créé `Docker Agent`, un framework agentique pour créer des agents.

Des agents de code, mais pas seulement.

`Open-Source, Model-Agnostic, souvent sans une ligne de code.`

---

# Qui sommes-nous ?

**Djordje Lukić** - Principal Engineer à **Docker**

**David Gageot** - Senior Principal Engineer à **Docker**

## Produits IA depuis 1 an et demi:

+ `Gordon`
+ `MCP Gateway`
+ `MCP Toolkit`
+ `Docker Agent`
+ Et plus 🐳

---

# L'IA depuis un an et demi

L'impression d'une éternité !

+ Débuts sur Gordon avec `llama3` puis `gpt-4o`
+ Pas de `MCP`
+ Pas de `Opus`
+ Débuts de `Copilot`
+ Des chat bots mais pas d'agents
+ Coder avec l'IA est possible mais une route semée d'embûches
+ Prémices du *vibe-coding*

---

# Et maintenant ?

Un monde nouveau:

+ `Opus` ~~4.6~~`4.7`, `GPT 5.4`
+ `MCP` est déjà presque out
+ Des **orchestrateurs** d'agents
+ Des codebase `100%` codées par des agents

---

# Et vous ?

+ Qui utilise l'IA pour coder ?

---

# Et vous ?

+ Qui utilise l'IA pour coder ?
+ Quels outils ?

---

# Et vous ?

+ Qui utilise l'IA pour coder ?
+ Quels outils ?
+ Qui ne code plus ?

---

# Et vous ?

+ Qui utilise l'IA pour coder ?
+ Quels outils ?
+ Qui ne code plus ?
+ Qui créé des agents pour autre chose que coder ?

---

# Docker Agent

## Un vrai Couteau Suisse
+ Open-Source
+ Ouvert: `OpenAI`, `Anthropic`, `Gemini`, `Mistral`, `Docker Model Runner`, `Ollama` ...

## Riche
+ `Docker Agent` est codé principalement avec `Docker Agent`
+ Beaucoup, beaucoup de fonctionnalités !

## Utilisable de plusieurs manières
+ `YAML`
+ `SDK` Go
+ Utilisé par `Docker` en production pour `Gordon`

---

# Première Demo

`Docker Agent` sert à définir tout type d'agent.

**Du plus utile au moins utile.**

## Exemple totalement inutile:

```bash
$ docker agent run "pirate.yaml"
```

---

# Démo - Pirate

## Du YAML et rien d'autre

```yaml
agents:
  root:
    model: anthropic/claude-sonnet-4-5
    instruction: Always answer by talking like a pirate.
    welcome_message: |
      Ahoy! I be yer pirate guide, ready to set sail on the seas o' knowledge!
      What be yer quest? 🏴‍☠️
```

---

# Démo - Agents de Code

`Docker Agent` est un excellent Agent de Code. Extrêmement flexible.

`Docker Agent` est codé principalement avec `Docker Agent`

+ Multi-agents
+ Multi-modèles
+ Skills
+ Built-in tools
+ MCP
+ AGENTS.md
+ Web search
+ RAG
+ ...

## Une simple commande

```bash
$ docker agent run "coder"
```

**Mais le mieux est de créer son propre agent.**

---

# Démo - Créer son Agent de Code

Définir un agent principal en `YAML`.

+ Ici, avec un modèle `Anthropic`
+ Peut être (`OpenAI`, `Gemini`, `Docker Model Runner`, `Ollama`, `Mistral`, `OpenAI compatible`...)
+ Tous les tools nécessaires à manipuler du code

```yaml
agents:
  root:
    model: anthropic/claude-sonnet-4-6
    instruction: |
      Help with code-related tasks by examining, modifying, and
      validating code changes...
    skills: true
    add_prompt_files:
      - AGENTS.md
    toolsets:
      - type: filesystem
      - type: shell
      - type: mcp
        command: gopls
        args: ["mcp"]
```

---

# Démo - Ajouter un "Planificateur"

Ajouter un sous-agent capable de préparer le travail en posant des questions
à l'utilisateur.

```yaml
agents:
  planner:
    model: anthropic/claude-opus-4-6
    instruction: |
      You are a planning agent responsible for gathering user requirements
      and creating a development plan.
      
      Always ask clarifying questions to ensure you fully understand
      the user's needs before creating the plan.
      Once you have a clear understanding, analyze the existing code
      and create a detailed development plan in a markdown file.
      Do not write any code yourself.

      Once the plan is created, you will delegate tasks to the root agent.
      Make sure to provide the file name of the plan when delegating.
      Write the plan in the current directory.

      Use the `user_prompt` tool to ask questions to the user.
      Prefer Multiple Choice Questions.
    toolsets:
      - type: filesystem
      - type: user_prompt
```

---

# Démo - Ajouter un "Bibliothécaire"

Pour faire des recherches Web ou obtenir des informations sur les APIs.

```yaml
models:
  gemini:
    provider: google
    model: gemini-3.1-flash-lite-preview
    provider_opts:
      google_search: true

agents:
  librarian:
    model: gemini
    instruction: |
      You are the librarian, your job is to look for relevant documentation
      to help the golang developer agent.

      When given a query, search the internet for relevant documentation,
      articles, or resources that can assist in completing the task.

      Use context7 for searching documentation.
    toolsets:
      - type: fetch
      - type: mcp
        ref: docker:context7
```

---

# Démo - Version finale

Nous avons notre propre version d'un *Agent de Code*.

## Une simple commande

```console
$ docker agent run "./coder.yaml"
```

---

# Démo - Partage

Cet agent est **facile à partager !** 

## Publier

```console
$ docker agent share push "./coder.yaml" "davidgageot135/coder:devoxx"
```

## Utiliser

```console
$ docker agent run "davidgageot135/coder:devoxx"
```

---

# Démo - Un agent spécialisé

On peut donc avoir un `Agent de Code` ultra-performant
et à la fois très flexible.

**Peut-on écrire des agents plus spécialisés ?**

+ Pouvant utiliser des modèles `moins couteux`
+ Et `plus rapides`
+ Donc avec un `choix plus large` de modèles
+ Avec le `minimum d'outils`, afin de limiter l'impact de possibles erreurs

**La réponse est oui, nous allons créer un agent de ce type.**

---

# Démo - Un agent spécialisé

Il y a des `APIs` tout autour de nous.

*Peut-on écrire un agent qui utilise ces APIs comme des outils ?*

## Transformer un Open API en boite à outils

```yaml
- type: openapi
  url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

**Ici, c'est une API Pokemons, ailleurs, c'est une API métier.**

---

# Démo - Un agent spécialisé en Pokemons


```yaml
agents:
  root:
    model: anthropic/claude-haiku-4-5
    instruction: |
      Tu es un expert en Pokémon.
      Sois bref dans tes réponses.
      Sois aussi amusant !
      IMPORTANT : Utilise toujours les noms français des Pokémon.
    toolsets:
      - type: openapi
        url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

## Lancement de l'agent

```bash
$ docker agent run "./pokemon.yaml"
```

---

# Démo - Un agent spécialisé en Pokemons


```yaml
agents:
  root:
    model: anthropic/claude-haiku-4-5
    instruction: |
      Tu es un expert en Pokémon.
      Sois bref dans tes réponses.
      Sois aussi amusant !
      IMPORTANT : Utilise toujours les noms français des Pokémon.
    toolsets:
      - type: openapi
        url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

## Lancement de l'agent

```bash
$ docker agent run "./pokemon.yaml"
```

## Ou une version plus avancée:

```bash
$ docker agent run "./pokemon-plus.yaml"
```

---

# Et après ?

Comment décider si cet agent est optimisé pour une liste de tâches ?

## Evals

On écrit des `evals` = des cas de test:

```yaml
{
  "title": "Pokémon aléatoire: informations complètes",
  "evals": {
    "relevance": [
      "Includes the url to this pokemon in the pokedex"
    ]
  },
  "messages": [
    {
      "message": {
        "message": {
          "role": "user",
          "content": "Choisis un Pokémon au hasard et dis-moi tout sur lui",
        }
...
```

---

# Démo - Comment optimiser cet agent ?

## Lancement des evals

```script
$ docker agent eval "./pokemon-plus.yaml"

Validating judge model...
Pre-building 1 Docker image(s)...
Running 2 evaluations with concurrency 16

✗ Pokémon aléatoire: informations complètes ($0.049806)
  ✓ tool calls
  ✗ relevance
✗ Pokémon de départ de Kanto - Classement stats ($0.135214)
  ✗ tool calls score 0.67
  ✗ relevance

✅     Tool Calls: 83.3% avg F1 (2 evals)
❌      Relevance: 1/3 passed (33.3%)

Total Cost: $0.185020
Total Time: 30s
Sessions DB: evals/results/warm-dawn-182.db
```

---

# Et ensuite ?

On modifie le `YAML` à la main et on boucle !

---

# Et ensuite ?

On modifie le `YAML` à la main et on boucle !

## Optimisation automatique

Bien mieux: On utilise un autre agent pour optimiser notre agent.

```bash
$ docker agent run "./eval-expert.yaml" "Optimize pokemon-plus.yaml"
```

---

# Et ensuite ?

On modifie le `YAML` à la main et on boucle !

## Optimisation automatique

Bien mieux: On utilise un autre agent pour optimiser notre agent.

```bash
$ docker agent run "./eval-expert.yaml" "Optimize pokemon-plus.yaml"
```

## Résultat

```diff
$ diff pokemon-plus.yaml pokemon-plus-optimized.yaml

>       Quand tu présentes des informations sur un Pokémon, inclus TOUJOURS l'URL vers sa page dans le Pokédex (ex: https://pokeapi.co/api/v2/pokemon/25/ pour Pikachu).
>
>       Quand tu compares plusieurs Pokémon dans un tableau récapitulatif, tu DOIS utiliser des emojis pour indiquer DEUX choses pour CHAQUE stat individuelle (PV, ATK, DEF, ATK Spé, DEF Spé, VIT) :
>       - 🟢 à côté de la valeur du Pokémon qui a la MEILLEURE valeur pour cette stat
>       - 🔴 à côté de la valeur du Pokémon qui a la PIRE valeur pour cette stat
>       Tu dois TOUJOURS marquer à la fois le meilleur ET le pire pour chaque stat. Ne marque jamais seulement le meilleur sans le pire.
```

---

# Go SDK

On peut aller encore plus loin en combinant du `Yaml`, du `code` et le `SDK` de `Docker Agent`.

## Un exemple

Un outil capable de lire des tickets de caisse et d'en extraire des montants.

```
~~~mermaid-ascii
graph LR
Image --> Go --> SDK Docker Agent --> Agent Yaml --> JSON
~~~
```

---

# Démo - Lecture de Tickets

```yaml
agents:
  root:
    model: openai/gpt-5.4
    instruction: |
      Your job is to read a receipt and extract the total price from it.
      You will be given the receipt as text.
      You should only return the total price in a structured format.
    ...
```

---

# Démo - Lecture de Tickets

```yaml
agents:
  root:
    ...
    structured_output:
      name: "ticket"
      description: "Informations extraites du ticket de caisse"
      strict: true
      schema:
        type: object
        properties:
          store:
            type: string
            description: "The store name"
          price:
            type: number
            description: "The total price of the purchase"
        required:
          - store
          - price
        additionalProperties: false
      ...
```

---

# Démo - Lecture de Tickets

```yaml
agents:
  root:
    ...
    toolsets:
      - type: filesystem
```

## Lancement

```bash
go run ./...
```

---

# On ne vous a pas montré

+ A2A
+ ACP
+ MCP (local, remote, oauth)
+ Handoff
+ Alloy
+ RAG
+ Code Mode
+ Slash commands
+ Hooks
+ Sandbox
+ Mode Routing
+ More tools (Http, TODO, Memory, Scripts, LSP, Background Agents...)
+ ...

---

# Questions

Merci à tous !
