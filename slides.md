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

# Qui sommes-nous?

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
+ Pas de `Claude Code`
+ Débuts de `Copilot`
+ Des chat bots mais pas d'agents
+ Coder avec l'IA est possible mais une route semée d'embûches
+ Prémices du *vibe-coding*

---

# Et maintenant?

Un monde nouveau:

+ `Claude Code`
+ `Opus` ~~4.6~~`4.7`, `GPT 5.4`
+ `MCP` est déjà presque out
+ Des **orchestrateurs** d'agents
+ Des codebase `100%` codées par des agents

---

# Et vous?

+ Qui utilise l'IA pour coder?

---

# Et vous?

+ Qui utilise l'IA pour coder?
+ Qui ne code plus?

---

# Et vous?

+ Qui utilise l'IA pour coder?
+ Qui ne code plus?
+ Qui créé des agents pour autre chose coder?

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
$ docker agent run "pirate"
```

---

# Demo: Pirate

## Du YAML et rien d'autre

```yaml
agents:
  root:
    model: auto
    instruction: Always answer by talking like a pirate.
    welcome_message: |
      Ahoy! I be yer pirate guide, ready to set sail on the seas o' knowledge!
      What be yer quest? 🏴‍☠️
```

---

# Demo - Coding Agent

`Docker Agent` est un excellent Coding Agent.

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

**Mais le mieux est de créer le sien.**

---

# Demo - Créer son Coding Agent

Définir un agent principal en `YAML`.

+ Ici, avec un modèle `Anthropic`
+ Peut être (`OpenAI`, `Gemini`, `Docker Model Runner`, `Ollama`, `Mistral`, `OpenAI compatible`...)
+ Tous les tools nécessaires à manipuler du code

## Le YAML

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
      - type: todo
      - type: mcp
        command: gopls
        args: ["mcp"]
```

---

# Demo - Ajouter un agent "Planificateur"

Ajouter un sous-agent capable de préparer le travail en posant des questions
à l'utilisateur.

## Le YAML

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

# Demo - Ajouter un agent de "Bibliothecaire"

Pour faire des recherches Web ou obtenir des informations sur les APIs.

## Le YAML

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

# Demo - Version finale

Nous avons notre propre version d'un *Agent de Code*.

## Lancement de l'agent

```console
$ docker agent run "./coder.yaml"
```

---

# Demo - Partage

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

# Demo - Un agent spécialisé

On peut donc avoir un `Agent de Code` ultra-performant
et à la fois très flexible.

*Peut-on écrire des  agents plus spécialisés ?*

+ Pouvant utiliser des modèles *moins couteux*
+ Et *plus rapides*
+ Donc avec un *choix plus large* de modèles
+ Avec le minimum d'outils, afin de limiter l'impact de possibles erreurs

**La réponse est oui, nous allons créer un agent de ce type.**

---

# Demo - Un agent spécialisé

Il y a des `APIs` tout autour de nous.

*Peut-on écrire un agent qui utilise ces APIs comme des outils ?*

Il est facile de transformer un API compatible Open API en une
boite à outils.

```yaml
- type: openapi
  url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

**Ici, c'est une API Pokemons, ailleurs, c'est une API métier.**

---

# Demo - Un agent spécialisé en Pokemons

## Le YAML

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

# Demo - Un agent spécialisé en Pokemons

## Le YAML

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

ou une version encore plus riche:

```bash
$ docker agent run "./pokemon-plus.yaml"
```

---

# Demo - Comment optimiser cet agent?

On écrit des `evals` et on vérifie si notre agent obtient un bon score.

## Lancement des evals

```bash
$ docker agent eval "./pokemon-plus.yaml"
```

## Et ensuite?

On modifie le `YAML` à la main et on boucle !

---

# Demo - Comment optimiser cet agent?

Bien mieux: Utiliser un autre agent pour optimiser notre agent.

## Optimisation automatique

```bash
$ docker agent run "./eval-expert.yaml" "Optimize pokemon-plus.yaml"
```

---

# Go SDK

On peut aller encore plus loin en combinant du `Yaml`, du `code` et le `SDK` de `Docker Agent`.

## Un exemple

Un outil capable de lire des tickets de caisse et d'en extraire des montant.

```
~~~mermaid-ascii
graph LR
Image --> Go --> SDK Docker Agent --> Agent Yaml --> JSON
~~~
```

---

# Demo - Lecture de Tickets

## Yaml

```yaml
agents:
  root:
    model: openai/gpt-5.4
    instruction: Your job is to read a receipt and extract the total price from it. You will be given the receipt as text. You should only return the total price in a structured format.
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
    toolsets:
      - type: filesystem
```

## Lancement

```bash
go run ./...
```

---

# Questions

Merci à tous !
