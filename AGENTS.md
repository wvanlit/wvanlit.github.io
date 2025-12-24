# Repository Guidelines

## Project Structure & Module Organization
This is an Astro-based portfolio + blog site with an editorial layout and a custom theme toggle.
- `src/` holds all application code. Add routes in `src/pages/` (e.g., `src/pages/index.astro`), and shared UI in `src/components/` if needed.
- `src/layouts/BaseLayout.astro` is the global shell. It renders the masthead (name, bio, nav), recent posts, and the theme toggle.
- `src/content/` stores Markdown collections for blog posts. The blog collection is `src/content/blog/`.
- `public/` is for static assets served as-is (images, fonts, robots.txt). The resume PDF should live at `public/resume.pdf`.
- `dist/` is the generated build output (do not edit by hand).
- Root config lives in `astro.config.mjs`, `tsconfig.json`, and `package.json`.

## Build, Test, and Development Commands
All commands run from the repo root:
- `pnpm install` installs dependencies using the checked-in `pnpm-lock.yaml`.
- `pnpm dev` starts the local dev server at `http://localhost:4321`.
- `pnpm build` creates a production build in `dist/`.
- `pnpm preview` serves the production build locally for verification.
- `pnpm astro -- --help` lists additional Astro CLI options.

## Coding Style & Naming Conventions
There is no custom formatter or linter configured. Keep formatting consistent with existing files and Astro defaults.
- Pages/routes: `kebab-case` (e.g., `src/pages/about-us.astro`).
- Components: `PascalCase` (e.g., `src/components/HeroBanner.astro`).
- Assets: match usage, prefer lowercase with dashes (e.g., `public/logo-mark.svg`).
- Use ASCII unless a file already contains non-ASCII characters.

## Content, Layout, and UX Guidelines
This site reads as a senior developer portfolio, resume, and blog. Keep copy concise, outcome-focused, and scannable.
- The masthead is in the main layout, so avoid duplicating global navigation or identity in pages.
- Recent posts are pulled from the blog collection in `src/content/blog/`.
- The theme toggle is controlled in `src/layouts/BaseLayout.astro`; keep it subtle and animated.
- Use an editorial minimal visual style with generous whitespace and strong hierarchy.
- Keep navigation tight: Home, Work, Writing, Resume, Contact.
- Store long-form content (blog posts, case studies) in Markdown; keep global sections hardcoded in Astro.

## About the Author
Known context to carry forward when writing copy or structuring content:
- Name: Wessel van Lit.
- Brand/domain: litdev (litdev.nl, litdev.bearblog.dev, GitHub `wvanlit`).
- Role framing: senior back-end software engineer.
- Domain focus: distributed systems, event-driven platforms, and C# microservices.
- Preferred direction: premium editorial minimal visual style.
- Preferred tone: professional, friendly, and approachable.
