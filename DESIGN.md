# Overview

A small, fast personal site for a senior engineer. The site will feel light and a bit hand‑drawn, but it stays clear and easy to read. It highlights a short bio, work history, writing, book notes, and WIP projects. It ships with light and dark mode.

---

## Goals

* Put a clear, short bio up front.
* Make reading pleasant: crisp type, calm spacing, high contrast.
* Keep the “hand‑drawn” vibe as an accent, never as decoration that hurts legibility.
* Show work history on a vertical timeline (Now → 2018).
* Host a blog for development notes and essays.
* Host book reviews and reading notes.
* Show projects in progress with short blurbs and links.
* Load fast on slow networks.

## Non‑Goals

* No complex CMS; content is markdown/MDX in Git.
* No heavy animation or 3D.
* No tracking beyond privacy‑friendly analytics.

## Audience

* Hiring managers, peers, and conference organizers.
* Engineers who want to read your notes.

## Brand & Voice

* Friendly, curious, and direct.
* Plain English. Short sentences. Active voice.
* Light playfulness in small details (underlines, doodle borders), not in the copy.

---

# Information Architecture

**Top‑level navigation**

* Home (`/`)
* Writing (`/writing`)
* Books (`/books`)
* Projects (`/projects`)
* Work (`/work`)
* About (`/about`)

**Secondary**

* Uses (`/uses`) — optional
* Colophon (`/colophon`) — optional

**Footer**

* Email, GitHub, LinkedIn, Mastodon/Bluesky (icons + text labels)

---

# Page Blueprints

## Home (`/`)

**Purpose:** Quick intro and clear routes to content.

**Blocks (top to bottom):**

1. **Hero bio**: avatar, name, 2–3 line bio, location. Inline links to Work, Writing, Projects.
2. **Now** (one line): current role and focus.
3. **Featured writing**: 1–3 cards.
4. **Latest posts**: list of 5 with dates.
5. **Projects (WIP)**: 2–4 tiles with status chips.
6. **Books**: latest 3 reviews/notes.
7. **Mini‑timeline**: last two roles; link to full Work page.

## Writing (`/writing`)

* Filterable list by tag.
* Each item: title, short deck, date, tags, read time.
* Search input (client‑side fuzzy search).

**Post template**

* Title, deck (1–2 lines), date, read time.
* TOC (auto‑generated for h2/h3) collapsible on mobile.
* Code blocks with copy‑button.
* Footnotes and callouts.
* Prev/next at bottom.

## Books (`/books`)

* Two tabs or filters: **Reviews** and **Notes**.
* Grid/list with cover, title, author, rating (optional), status (Reading / Finished / Shelved).
* Book page: summary, highlights, your notes (MDX), link to external references.

## Projects (`/projects`)

* Grid of cards (image or emoji), short description, stack chips, status (Idea / WIP / Paused / Shipped), link(s).
* Project page: Problem, Approach, What’s next.

## Work (`/work`)

* Vertical timeline from **Now → 2018**.
* Each stop has: role, company, dates, 2–4 bullets, links.
* Uses `<time>` elements and an accessible ordered list.

## About (`/about`)

* Friendly photo.
* 4–6 short paragraphs: who you are, what you like to build, how to contact you.
* Optionally a short “values” list.

---

# Visual Design

## Layout & Rhythm

* Content width: 72ch max for prose.
* Grid: 12‑col fluid; cards snap to 4/3/2 across as space allows.
* Vertical rhythm: base unit 8px (use multiples).

## Typography

* **Body**: A humanist sans (e.g., Inter, IBM Plex Sans) at 17–18px; 1.6 line‑height.
* **Headings**: Same family with stronger weight; avoid huge sizes.
* **Monospace**: For code and inline tokens.
* **Hand‑drawn accent**: A subtle script/hand font used sparingly for small labels or headers (e.g., section doodle captions). Never for body text.

## Hand‑Drawn Accents (Tasteful)

* Doodle underline for links on hover.
* Rough “bracket” highlights for key callouts.
* Squircle/rough border around avatar.
* Occasional margin doodles on dividers.
* Keep stroke widths light; respect contrast in both themes.

## Icons & Illustration

* Simple strokes. No heavy fills. Prefer SVG.
* Use the same stroke width as doodles to keep harmony.

---

# Components & Patterns

## Global Header

* Left: “litdev” wordmark.
* Right: nav links, theme toggle.
* Mobile: menu button that reveals a sheet with the same links.

## Footer

* Social links, small copyright.

## Bio Hero

* Avatar with rough border.
* Name + short bio (2–3 lines).
* CTA buttons: Writing, Projects, Work.

## Timeline

* Semantic ordered list (`<ol>`).
* Each item: dot + vertical line, role, company, dates, bullets.
* Reduced motion mode removes any animation.

## Article

* Narrow measure, generous line‑height.
* TOC at top (collapsible on mobile).
* Pull‑quote and callout components.
* Code blocks with copy buttons and highlighted lines.

## Card

* Soft border, shadow at rest; lift on hover.
* Small doodle accent on hover (e.g., underline or sparkle) — optional.

---

# Accessibility

* AA contrast or higher across themes.
* Respect `prefers-reduced-motion`.
* Skip‑to‑content link.
* Keyboard focus outlines always visible; distinct from hover.
* Alt text on all images; figure/figcaption for book covers and screenshots.
* Semantic landmarks: `<header>`, `<main>`, `<nav>`, `<footer>`.

---

# Performance

* First load targets: LCP < 2.0s on 3G, CLS < 0.1, TBT < 200ms.
* Static generation for all content pages.
* Local fonts with `display: swap`; preload headers.
* Optimized images via Astro's Image component (`astro:assets`) with AVIF/WebP; set `sizes` for responsive images.
* Inline critical CSS; lazy‑load non‑critical images.
* Zero client JS on article pages (except theme + optional TOC); use Astro Islands to hydrate only where needed.

---

# Content Model (MD/MDX)

## Post (`/writing/[slug]`)

Frontmatter

```
title: ""
date: YYYY-MM-DD
description: ""
readingTime: 6
featured: true | false
tags: ["architecture", "react"]
hero:
  src:
  alt:
```

Body supports MDX (diagrams, callouts).

## Book (`/books/[slug]`)

```
title: ""
author: ""
status: reading | finished | shelved
rating: 1-5 # optional
year: 2025
cover:
  src:
  alt:
links:
  goodreads:
```

Body: review + notes sections.

## Project (`/projects/[slug]`)

```
title: ""
stack: ["ts", "react", "aws"]
status: idea | wip | paused | shipped
repo:
link:
```

Body: problem, approach, next.

## Work Item

```
year: 2025-2022
role: "Software Engineer"
company: "Coolblue"
bullets:
  - "…"
links:
  - label: "Project"
    href: ""
```

---

# Tech Stack

* **Framework:** Astro.
* **Styling:** Tailwind CSS + CSS variables for tokens.
* **Theme:** CSS `prefers-color-scheme` + a small Astro island for the toggle (no `next-themes`).
* **Content:** Astro Content Collections + `@astrojs/mdx` with Zod schemas for typed content.
* **Code Blocks:** Shiki via Astro's Markdown config (shikiji); add a copy button. Optional: `rehype-pretty-code` if needed.
* **Images:** Astro Image (`astro:assets`) with responsive `sizes` and optional blurred/dominant‑color placeholders.
* **Search:** Client‑side Fuse.js on `/writing` only.
* **Analytics:** Plausible (self‑host or cloud).
* **Hosting:** GitHub Pages (`wvanlit.github.io`) with custom domain `litdev.nl` (or any static host like Vercel/Netlify).

---

# Hand‑Drawn Implementation Notes

* Use a rough‑annotation library for underlines, boxes, and highlights on key phrases and section headings; keep stroke thin and short durations.
* Use a rough/sketch library for avatar border, dividers, and small doodles. Render to SVG for crispness.
* Provide a “No doodles” toggle (optional) if you want an extra‑clean mode.

---

# Light/Dark Mode

* Default to system setting; toggle persists in `localStorage`.
* Swap CSS custom properties per theme.
* Test contrast of accent/support colors in both modes.
* Images: prefer neutral backgrounds so avatars and book covers look good in both themes.

---

# Navigation & URL Conventions

* Post slugs are kebab‑case of title (`/writing/distributed-architecture-for-product-owners`).
* Year archives: `/writing/2025`.
* Tag pages: `/tag/[tag]` (optional; generate statically).

---

# Build Plan (3 short phases)

## Phase 1 — Foundations (1–2 days)

* Create Astro app; add `@astrojs/mdx` and set up Content Collections (posts, books, projects, work).
* Set up Tailwind, theme tokens, and light/dark mode.
* Build global layout, header, footer, basic pages.

## Phase 2 — Content & Components (2–3 days)

* Home hero, cards, lists.
* Article template with TOC and code.
* Timeline component.
* Books grid + detail.
* Projects grid + detail.

## Phase 3 — Polish (1–2 days)

* Hand‑drawn accents.
* Search on `/writing`.
* Open Graph images per page.
* Analytics, sitemap, redirects.

---

# Copy Drafts

**Home hero**

> Hi, I’m Wessel — a software engineer in Rotterdam. I build reliable systems with C#, TypeScript, and React. I like shipping small tools, writing about architecture, and reading good books.

**About intro**

> I’m a senior software engineer at Coolblue. I enjoy clear code, strong teams, and simple tools. This site is my notebook for work, books, and side projects.

---

# Wireframe Notes (text)

* **Header:** wordmark left, nav center/right, theme toggle right.
* **Home:** two‑column at desktop: bio left, featured writing right; stacks on mobile.
* **Writing:** filter/search at top; list of posts; sticky sidebar TOC on post pages.
* **Books:** grid of covers; filter chips; detail page with notes.
* **Projects:** cards with status chips; detail page.
* **Work:** timeline down the page; years as section headers.

---

# QA Checklist

* Keyboard‑only nav works.
* Light/dark pass contrast rules.
* Images have alt.
* Code blocks wrap and scroll on small screens.
* 404/500 pages exist.

---

# Nice‑to‑Haves (later)

* “/now” page.
* Digital garden style sub‑notes (unlisted).
* Email newsletter (Buttondown).
* Print CSS for posts.

---

# Out of Scope (v1)

* Comments system.
* Multi‑language.

---

# Deliverables

* Source repo with README and scripts.
* Content starter set: 1 featured post, 3 posts, 5 books, 3 projects, full timeline.
* Favicon, wordmark SVG, Open Graph defaults.

---

# Acceptance Criteria

* Deployed at `https://litdev.nl` with HTTPS.
* All pages pass Lighthouse Performance ≥ 90, Accessibility ≥ 95.
* Timeline lists roles from Now to 2018.
* System light/dark with a toggle.
* No layout shift on first load.
