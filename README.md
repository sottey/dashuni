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

First, clone the repository:

```bash
git clone https://github.com/sottey/dashuni.git
cd dashuni
```

Build it:

```bash
go build -o dashuni
```

You can now run `./dashuni` locally.

---

## ✅ Usage

```
dashuni [command] [flags]
```

### 📌 Commands

#### 1️⃣ convert
Render your universal JSON schema using a mapping template.

```
dashuni convert --input sample.json --mapping mappings/dashy.tmpl --output dashy-config.yml
```

**Flags:**
- `--input, -i` : Path to your universal site JSON
- `--mapping, -m` : Path to the Go text/template mapping file
- `--output, -o` : Where to write the rendered config

---

#### 2️⃣ validate
Check that your universal JSON schema is valid.

```
dashuni validate --input sample.json
```

**Flags:**
- `--input, -i` : Path to your universal site JSON

---

#### 3️⃣ list *(planned)*
List available mapping templates.

```
dashuni list
```

*(Future version)*

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

Examples you might include:

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

sottey  
[sottey on GitHub](https://github.com/sottey)

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

---

## ✅ Roadmap

✅ Validate universal JSON  
✅ Mapping templates for:
- Dashy
- Homer
- Honey
- LabDash
- Mafl (planned)
- Starbase (planned)

✅ CLI `list` command (planned)  
✅ Template variables and helpers

---

## ❤️ Why?

Because no one wants to manually recreate dashboards every time they switch apps.  
**Define once. Convert anywhere.**