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
+ Pas de `Opus`
+ Débuts de `Copilot`
+ Des chat bots mais pas d'agents
+ Coder avec l'IA est possible mais une route semée d'embûches
+ Prémices du *vibe-coding*

---

# Et maintenant?

Un monde nouveau:

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
+ Qui créé des agents pour autre chose que coder?

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
    model: anthropic/claude-sonnet-4-5
    instruction: Always answer by talking like a pirate.
    welcome_message: |
      Ahoy! I be yer pirate guide, ready to set sail on the seas o' knowledge!
      What be yer quest? 🏴‍☠️
```

---

# Demo - Un agent spécialisé

**Peut-on écrire des agents très spécialisés ?**

+ Pouvant utiliser des modèles *moins couteux*
+ Et *plus rapides*
+ Donc avec un *choix plus large* de modèles
+ Avec le minimum d'outils, afin de limiter l'impact de possibles erreurs

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

## Ou une version plus avancée:

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

# On ne vous a pas montré

+ Agents de Code
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
