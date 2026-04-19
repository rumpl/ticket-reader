---
author: David Gageot & Djordje Lukić
paging: Slide %d / %d
---

# Docker Agent, le Couteau Suisse Agentique de Docker.

L’IA révolutionne nos métiers autour du développement.
Chez Docker, nous jouons les éclaireurs et changeons totalement
notre façon de travailler. 

Pour ce faire, nous avons créé `Docker Agent`. un framework agentique pour
créer vos propres agents.

Des agents de code, mais pas seulement.

Le tout Open-Source, Model-Agnostic, souvent sans une ligne de code.

---

# Qui sommes-nous?

**Djordje Lukić** - Principal Engineer - Docker

**David Gageot** - Senior Principal Engineer - Docker

Produits IA depuis 1 an et demi

+ Gordon
+ MCP Gateway
+ MCP Toolkit
+ Docker Agent
+ Et plus 🐳

---

# L'IA depuis un an et demi

L'impression d'une éternité !

+ Débuts sur Gordon avec llama3 puis gpt-4o
+ Pas de MCP
+ Pas de Claude Code
+ Débuts de Copilot
+ Des chat bots mais pas d'agents
+ Coder avec l'IA est possible mais une route semée d'embûches
+ Prémices du vibe-coding

---

# Et maintenant?

Un monde nouveau.

+ Des agents partout
+ Claude Code
+ Opus 4.6, GPT 5.4
+ MCP est déjà presque out
+ Des orchestrateurs d'agents
+ Des codebase 100% codées par des agents

---

# Docker Agent

+ Un Couteau Suisse
+ Open-Source
+ Ouvert: OpenAI, Anthropic, Google, Mistral, DMR, ...
+ `Docker Agent` est codé principalement avec `Docker Agent`
+ Utilisable sans une ligne de code
+ Utilisé par Docker en production pour Gordon

---

# Demo - Pirate

`Docker Agent` sert à définir tout type d'agent.

+ Sans une ligne de code
+ Du plus utile au moins utile.

Un exemple totalement inutile:

```bash
$ docker agent run pirate
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

```bash
$ docker agent run coder
```

Mais vous devriez créer le votre.

---

# Demo - Créer son Coding Agent

Définir un agent principal en YAML.

+ Ici, avec un modèle Anthropic
+ Peut être (OpenAI, Gemini, Docker Model Runner, Ollama, Mistral, OpenAI compatible...)
+ Tous les tools nécessaires

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

```bash
$ docker agent run ./coder.yaml
```

---

# Demo - Un agent spécialisé

On peut donc avoir un Agent de Code ultra performant
et à la fois très flexible.

Peut-on écrire un agent beaucoup plus spécialisé ?

+ Pouvant utiliser des modèles moins couteux
+ Et plus rapides
+ Donc avec un choix plus large de modèles
+ Avec le minimum d'outils, afin de limiter l'impact de possibles erreurs

La réponse est oui, nous allons créer un agent de ce type.

---

# Demo - Un agent spécialisé

Il y a des APIs tout autour de nous.

Peut-on écrire un agent qui utilise ces APIs comme des outils ?

Il est facile de transformer un API compatible Open API en une
boite à outils.

```yaml
- type: openapi
  url: https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml
```

Ici, c'est une API Pokemons, ailleurs, c'est une API métier.

---

# Demo - Un agent spécialisé en Pokemons

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
        url: https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml
```

Ici, c'est une API Pokemon, ailleurs, c'est une API métier.

```bash
$ docker agent run ./pokemon.yaml
```

ou

```bash
$ docker agent run ./pokemon-plus.yaml
```

---

# Demo - How do you optimize this agent?

```bash
$ docker agent eval ./pokemon-plus.yaml
```

And manually make the changes...

---

# Demo - How do you optimize this agent?

Way better: User another agent to optimize it:

```bash
$ docker agent run ./eval-expert.yaml "Optimize pokemon-plus.yaml"
```

---

# Demo - Go SDK

---

# Questions

Q/A
