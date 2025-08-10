## AI guide for this repo (Astro personal site)

Purpose: help AI agents work productively in this codebase. Keep changes incremental, align with the product spec in `DESIGN.md`, and favor simple, static output unless interactivity is explicitly needed.

Big picture
- Framework: Astro (static site). No JS frameworks are installed; default is zero client JS.
- Current state: minimal starter with one page at `src/pages/index.astro`. Use file‑based routing to add pages.
- Product spec & IA: `DESIGN.md` is the source of truth for navigation, content types, visual tone, and performance/accessibility principles.

Run/build
- Scripts (see `package.json`): `npm run dev`, `npm run build`, `npm run preview`. Output builds to `dist/`.
- Dev server path: Astro default (see terminal output for the port). Hot reload is enabled.

Repo layout
- `src/pages/**/*.astro|md`: routes. Example: `src/pages/about.astro` => `/about`.
- `public/`: static assets served at site root (e.g., `public/favicon.svg` -> `/favicon.svg`).
- `astro.config.mjs`: Astro config (currently default). Update here when adding integrations.
- `tsconfig.json`: extends `astro/tsconfigs/strict` for stronger checks.

Conventions
- Prefer static rendering. Only add client interactivity via Astro Islands when strictly needed by the spec.
- Keep styles lightweight. If you introduce a styling system (e.g., CSS variables or a utility framework), centralize tokens and follow the color/spacing guidance in `DESIGN.md`.
- Routes and slugs must match the IA in `DESIGN.md` (Home, Writing, Books, Projects, Work, About). Use kebab‑case for slugs.
- Follow accessibility and performance notes in `DESIGN.md`: high contrast, semantic landmarks, reduced‑motion support.

When adding pages/content
- Simple page: create `src/pages/<name>.astro` with semantic HTML. Example: `src/pages/about.astro` for `/about`.
- Section index pages: `src/pages/writing/index.astro`, `src/pages/projects/index.astro`, etc.
- Assets: place images in `public/` and reference by absolute path (`/images/cover.jpg`).

Examples
- New route: add `src/pages/work.astro` with an ordered list (`<ol>`) timeline per `DESIGN.md`’s Work blueprint.
- Home composition: implement blocks in order (hero bio, featured writing, latest posts, projects, books, mini‑timeline). Use simple server‑rendered markup first.

Extending the system (only if asked)
- Content types (posts/books/projects) and MD/MDX: introduce Astro Content Collections and schemas to match the frontmatter examples in `DESIGN.md`. Place content under `src/content/<collection>/...` and wire list/detail pages accordingly.
- Theming: add CSS custom properties for light/dark and a minimal theme toggle island. Keep non‑article pages as close to zero JS as possible.

Quality bar
- Ship pages with semantic HTML, accessible focus, and no layout shift. Optimize images with widths/sizes if you add Astro’s image tooling later.
- Avoid adding heavy dependencies without a clear need tied to `DESIGN.md`.

References
- Spec: `DESIGN.md`
- Starter docs: `README.md`

Questions for refinements
- Do you want Content Collections wired now, or keep pages hardcoded until content is ready?
- Should we introduce a base layout and shared header/footer next, or build individual pages first?
