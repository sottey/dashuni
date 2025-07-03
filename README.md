# 🎨 Dashuni

Welcome to **Dashuni** — your universal *dashboard config converter* for the homelab world!

Ever switch from **Dashy** to **Homer** to **Honey** to **Starbase** and think:

> *“I really don’t want to re-enter my 50 services again…”*

Dashuni *solves* that. You keep one single source of truth about your servers, services, and bookmarks, and Dashuni generates the correct config for your favorite dashboard.

---

## 🚀 What is Dashuni?

Dashuni is a command-line tool written in Go. It takes a **universal JSON** describing your homelab layout and services, then uses **Go templates** to convert that into the native config formats for dozens of popular dashboards.

✅ One source of truth.  
✅ Export to anything.  
✅ No more tedious re-entry.

---

## ‼️ What is Dashuni NOT?

Dashuni is NOT for ALL dashboards. Because it focuses on writing config files in YAML, JSON, etc., it is not (yet) able to handle dashboards which use MongoDB, Postgres, or another storage mechanism that is not a plain text configuration.

---

## ⭐️ Key Features

✨ Convert your homelab service list to **many dashboards**  
✨ Templated and **extensible**  
✨ Supports **icon mapping** (for Font Awesome, MDI, Simple Icons)  
✨ Easy **CLI** usage  
✨ Fast and scriptable  
✨ Bring your *own templates* to add new dashboards!

---

## 📦 Supported Dashboards

Out of the box, Dashuni can generate configs for:

- [**Dashy**](https://github.com/Lissy93/dashy)
- [**Homer**](https://github.com/bastienwirtz/homer)  
- [**Honey**](https://github.com/dani3l0/honey)  
- [**LabDash**](https://github.com/AnthonyGress/lab-dash)  
- [**MAFL**](https://github.com/hywax/mafl)  
- [**Starbase**](https://github.com/notclickable-jordan/starbase-80)  
- [**Hiccup**](https://github.com/ashwin-pc/hiccup)  
- [**GetHomepage**](https://github.com/gethomepage/homepage)

And it’s easy to add more!

---

## 💻 How Does it Work?

1️⃣ You create a **universal JSON** describing your services in a normalized schema.  
2️⃣ You pick a template for your target dashboard.  
3️⃣ Dashuni renders the template with your data → you get the correct config format.  
4️⃣ Paste it into your dashboard and enjoy.

✅ No more manual re-entry.  
✅ Consistency between all your dashboards.

---

## 🗂️ Example Universal Schema

Here’s a snippet of your **one source of truth**:

```json
{
  "site": {
    "name": "My Homelab",
    "description": "Central dashboard for all services",
    "theme": "auto",
    "pages": [
      {
        "title": "Main",
        "sections": [
          {
            "title": "Servers",
            "items": [
              {
                "title": "Dockge",
                "url": "http://192.168.7.2:8080",
                "icon": "https://example.com/icons/dockge.png",
                "description": "Docker management UI",
                "tags": ["docker", "admin"],
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

You maintain this single file. Dashuni does the rest.

---

## ⚙️ Installation

Clone the repo:

```bash
git clone https://github.com/sottey/dashuni.git
cd dashuni
```

Build the binary:

```bash
go build -o dashuni
```

Or install via Go:

```bash
go install github.com/sottey/dashuni@latest
```

---

## 🧰 CLI Usage

To convert:

```bash
./dashuni convert --input mylab.json --mapping mappings/dashy.tmpl --output dashy-config.yml
```

Options:

- `--input`: your universal JSON
- `--mapping`: target dashboard template
- `--output`: where to save the result

List your available templates:

```bash
./dashuni list
```

---

## 🎨 Templating System

Dashuni is powered by **Go’s text/template** system.

Your template defines exactly how your universal data maps to the dashboard config.

**Example Dashy Template Snippet:**

```gotemplate
appConfig:
  theme: "{{ .Site.Theme }}"
sections:
{{- range .Site.Pages }}
  {{- range .Sections }}
  - name: "{{ .Title }}"
    items:
      {{- range .Items }}
      - title: "{{ .Title }}"
        url: "{{ .URL }}"
        icon: "{{ .Icon }}"
      {{- end }}
  {{- end }}
{{- end }}
```

✅ Fully customizable  
✅ Add your own templates for new dashboards!

---

## 🗺️ Icon Mapping

Some dashboards don’t use raw URLs for icons but want standardized names like **mdi:docker** or **simple-icons:vaultwarden**.

✅ Dashuni supports an optional **fa-mapping.json**:

```json
{
  "Dockge": "mdi:docker",
  "FreshRSS": "mdi:rss",
  "VaultWarden": "simple-icons:vaultwarden"
}
```

Your template can use this to transform service names into correct icon identifiers.

---

## 📌 Note on Icons

**Important:** Not all icon libraries have all services!  

For example:

- **Simple Icons** only includes popular brands.
- Self-hosted projects (like Dockge or FreshRSS) often lack official icons.

You might need to:

- Use **mdi** alternatives
- Upload custom icons to your dashboard
- Maintain your own mapping JSON

---

## ⚡️ Adding New Dashboards

✅ Want to support another dashboard?  

Just:

1️⃣ Study the dashboard’s config format  
2️⃣ Write a new Go template in `mappings/`  
3️⃣ Use Dashuni to render your universal JSON into the new format

No code changes required. Templates are *data-driven*.

---

## 🏗️ Building Dashuni

```bash
go build -o dashuni
```

Check:

```bash
./dashuni --help
```

---

## 🤝 Contributions Welcome

Want to help make Dashuni better?

✅ Add new dashboard templates  
✅ Improve existing mappings  
✅ Fix bugs  
✅ Enhance CLI UX  
✅ Add tests

---

### To contribute:

1. Fork the repo  
2. Make your changes in a new branch  
3. Submit a pull request  

Let’s make dashboard switching *painless* for everyone!

---

## 🙏 Credits

Dashuni is designed and maintained by **sottey**, with friendly help from ChatGPT (who wrote this readme!).

Thanks to:

- All the homelab dashboard developers
- Go’s powerful `text/template` library
- Open source communities everywhere

---

## 🪄 License

MIT. See [LICENSE](./LICENSE).

---

> ✨ *Stop rewriting configs. Start enjoying your homelab.*  
> **Happy dashboarding! 🚀**