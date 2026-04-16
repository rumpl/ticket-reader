---
author: David Gageot & Djordje Lukic
paging: Slide %d / %d
---

# Docker Agent, le couteau suisse agentique de Docker.

L’IA révolutionne nos métiers autour du développement.
Chez Docker, nous jouons les éclaireurs et changeons totalement
notre façon de travailler. 

Pour ce faire, nous avons créé `Docker Agent`. un framework agentique qui
vous aide à créer vos propres agents. Des agents de code, mais pas seulement.

Le tout en open source, model-agnostic, souvent sans une ligne de code.

---

# Qui sommes-nous?

**Djordje Lukić** - Principal Engineer à Docker

**David Gageot** - Senior Principal Engineer à Docker

Sujets IA depuis 1 an et demi

+ Gordon
+ MCP Gateway
+ MCP Toolkit
+ Docker Agent
+ Et plus 🐳

---

# L'IA depuis un an et demi

L'impression d'une éternité.

+ Débuts sur Gordon avec llama3.1 puis gpt-4o
+ Pas de MCP
+ Pas de Claude
+ Copilot
+ Des chat bots mais pas d'agents
+ Coder avec l'IA est possible mais route semée d'embuches
+ Prémices du vibe-coding

---

# Maintenant

Un monde nouveau.

+ Des agents partout
+ Claude Code
+ Opus 4.6, GPT 5.4
+ MCP est déjà presque out
+ Des orchestrateurs d'agents
+ Des codebase 100% codées par des agents

---

# Docker Agent

+ Un couteau suisse
+ Ouvert (OpenAI, Anthropic, Google, Mistral, DMR, ...)
+ Open-source
+ Docker Agent est codé principalement avec Docker Agent
+ Utilisable sans une ligne de code
+ Utilisé en production pour Gordon

---

# Demo - Pirate

Docker Agent sert à définir tout type d'agent.

+ Sans une ligne de code
+ Du plus utile au moins utile.

Un exemple inutile:

```bash
docker agent run pirate
```
---

# Demo - Coding Agent

Docker Agent est aussi un excellent Coding Agent.

+ Multi-agents
+ Multi-model
+ Skills
+ Built-in tools
+ MCP
+ AGENTS.md
+ Web search
+ ...

```bash
docker agent run coder
```

Mais vous devriez créer le votre.

---

# Demo - Créer son Coding Agent

Définir un agent principal en YAML.

- Ici, avec un modèle Anthropic mais peut être (OpenAI, Gemini, DMR, Ollama, Mistral, OpenAI compatible...)
- Tous les tools nécessaires

```yaml
models:
  sonnet:
    provider: anthropic
    model: claude-sonnet-4-6

agents:
  root:
    model: sonnet
    instruction: Help with code-related tasks by examining, modifying, and validating code changes...
    skills: true
    add_environment_info: true
    add_prompt_files:
      - AGENTS.md
    sub_agents:
      - librarian
    toolsets:
      - type: filesystem
      - type: shell
      - type: todo
      - type: mcp
        command: gopls
        args: ["mcp"]
        version: "golang/tools@v0.21.0"
        tools: ["go_workspace", "go_symbol_references", "go_search", "go_rename_symbol", "go_package_api", "go_file_context"]
```

---

# Demo - Ajouter un agent "Planificateur"

```yaml
models:
  opus:
    provider: anthropic
    model: claude-opus-4-6
    thinking_budget: adaptive/low

agents:
  planner:
    model: opus
    instruction: |
      You are a planning agent responsible for gathering user requirements and creating a development plan.
      Always ask clarifying questions to ensure you fully understand the user's needs before creating the plan.
      Once you have a clear understanding, analyze the existing code and create a detailed development plan in a markdown file. Do not write any code yourself.
      Once the plan is created, you will delegate tasks to the root agent. Make sure to provide the file name of the plan when delegating. Write the plan in the current directory.
      Use the `user_prompt` tool to ask questions to the user. Prefer Multiple Choice Questions.
    toolsets:
      - type: filesystem
      - type: user_prompt
    sub_agents:
      - root
```

---

# Demo - Ajouter un agent de "Bibliothecaire"

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
      You are the librarian, your job is to look for relevant documentation to help the golang developer agent.
      When given a query, search the internet for relevant documentation, articles, or resources that can assist in completing the task.
      Use context7 for searching documentation.
    toolsets:
      - type: fetch
      - type: mcp
        ref: docker:context7
```
