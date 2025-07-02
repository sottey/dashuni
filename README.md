# 🧭 dashuni

> **dashuni** is a CLI tool written in Go to help you *convert a universal homelab dashboard schema* into the config formats of different dashboard apps.

No more retyping your entire server and service list when you switch dashboards!  

---

## ⭐️ Features

✅ Define your servers, services, and bookmarks in a single **universal JSON schema**  
✅ Convert to multiple dashboards using **mapping templates**  
✅ Add new dashboards just by adding a template file—no recompiling required!  
✅ Simple, scriptable CLI  

---

## 🚀 Installation

Clone the repository:

```bash
git clone https://github.com/sottey/dashuni.git
cd dashuni
```

Build it:

```bash
go build -o dashuni
```

Now you can run:

```bash
./dashuni
```

---

## ✅ Usage

```
dashuni [command] [flags]
```

### 📌 Commands

#### 1️⃣ convert

Render your universal JSON schema using a mapping template:

```
dashuni convert --input sample.json --mapping mappings/dashy.tmpl --output dashy-config.yml
```

**Flags:**
- `--input, -i` : Path to your universal site JSON
- `--mapping, -m` : Path to the Go text/template mapping file
- `--output, -o` : Path to write the rendered config

---

#### 2️⃣ validate

Check that your universal JSON schema is valid:

```
dashuni validate --input sample.json
```

**Flags:**
- `--input, -i` : Path to your universal site JSON

---

#### 3️⃣ list

List available mapping templates:

```
dashuni list
```

---

## ✅ Example Universal Schema

Example `sample.json`:

```json
{
  "site": {
    "name": "My Home Lab Dashboard",
    "description": "A dashboard for managing my home lab services",
    "favicon": "https://example.com/favicon.ico",
    "theme": "auto",
    "baseURL": "https://mydashboard.local",
    "version": "1.0.0",
    "pages": [
      {
        "title": "Main",
        "sections": [
          {
            "title": "Grump",
            "icon": "https://example.com/icons/grump.png",
            "items": [
              {
                "title": "Dockge",
                "url": "http://192.168.7.212:8081",
                "icon": "https://exmaple.com/icons/dockge.png",
                "description": "Dockge on Grump",
                "target": "_blank"
              }
            ]
          }
        ]
      }
    ]
  }
}
```

---

## ✅ Mapping Templates

Mapping templates live in the `mappings/` folder. They are standard Go `text/template` files.

Example structure:

```
mappings/
  dashy.tmpl
  homer.tmpl
  honey.tmpl
  labdash.tmpl
  mafl.tmpl
  starbase.tmpl
```

Each template converts the universal schema into the specific dashboard’s config format.

---

## ✅ Adding New Dashboards

To add a new target dashboard:

1. Create a new Go text/template file in `mappings/`.
2. Write your mapping using the universal schema model.
3. No recompilation needed—just use:

```
dashuni convert --input your.json --mapping mappings/new.tmpl --output new-config.yml
```

---

## ✅ Example Templates

✅ Dashy:
- sections with items
- supports status check

✅ Homer:
- sections with items
- subtitle field

✅ Honey:
- flat list of services
- no sections

✅ LabDash:
- page layout with desktop grid

✅ Mafl:
- YAML map of sections with service lists

✅ Starbase:
- JSON meta with flat links array

---

## ✅ Contributing

PRs welcome!

✅ Add new templates  
✅ Improve CLI  
✅ Add features

---

## ✅ License

MIT

---

## ⭐️ Author
[sottey on GitHub](https://github.com/sottey/dashuni)


---

## ✅ Quick Start

✔️ Build:

```bash
go build -o dashuni
```

✔️ Validate your universal config:

```bash
./dashuni validate --input sample.json
```

✔️ Convert to Dashy:

```bash
./dashuni convert --input sample.json --mapping mappings/dashy.tmpl --output dashy-config.yml
```

✔️ Convert to Homer:

```bash
./dashuni convert --input sample.json --mapping mappings/homer.tmpl --output homer-config.yml
```

✔️ List available templates:

```bash
./dashuni list
```

---

## ✅ Roadmap

✅ Validate universal JSON  
✅ Mapping templates for:
  - Dashy
  - Homer
  - Honey
  - LabDash
  - Mafl
  - Starbase

✅ CLI `list` command  
✅ Future ideas:
- More template variables and helpers
- Remote template repo support

---

## ❤️ Why?

Because no one wants to manually recreate dashboards every time they switch apps.  
**Define once. Convert anywhere.**