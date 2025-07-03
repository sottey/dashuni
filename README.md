# Dashuni

**Dashuni** is a universal converter for homelab dashboard configs.  

You define your homelab data in a single, simple JSON schema, then use Dashuni to render config files for any supported dashboard by applying a Go template.  

---

## ⭐️ Features

✅ Single source of truth for your homelab inventory  
✅ Supports multiple dashboard formats (Dashy, Homer, Honey, Labdash, Starbase, MAFL, etc.)  
✅ Clean, flexible Go templates  
✅ Optional Font Awesome icon mapping for dashboards that need it  
✅ Fully CLI-driven  

---

## 🚀 Installation

```bash
git clone https://github.com/sottey/dashuni.git
cd dashuni
go build -o dashuni
```

---

## ✅ Usage

```bash
./dashuni convert --input sample.json --mapping mappings/dashy.tmpl --output dashy-config.yml
```

**Options:**
- `--input` : your universal JSON description of the homelab
- `--mapping` : the Go text/template file to convert to your dashboard's config format
- `--output` : file to write the converted config to

---

## 📌 Input Schema

Your universal JSON should look like this:

```json
{
  "site": {
    "name": "My Home Lab Dashboard",
    "description": "A dashboard for managing my home lab services",
    "favicon": "https://example.com/favicon.ico",
    "theme": "auto",
    "version": "1.0.0",
    "pages": [
      {
        "title": "Infrastructure",
        "sections": [
          {
            "title": "Monitoring",
            "items": [
              {
                "title": "Grafana",
                "description": "Metrics and dashboards",
                "url": "https://grafana.local",
                "pingURL": "https://grafana.local/api/health",
                "icon": "https://example.com/icons/grafana.png",
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

## ✅ Template Data

Every template is executed with **this data**:

```
.
├── Site  (*model.Site)
│   ├── Name
│   ├── Description
│   ├── Favicon
│   ├── Theme
│   ├── Pages
│       ├── Sections
│           ├── Items
└── FAMap (optional map[string]string)
```

✅ Templates must use `.Site` prefix:

```gotemplate
.Site.Name
.Site.Pages
```

✅ Example:

```gotemplate
title: "{{ .Site.Name }}"
theme: "{{ .Site.Theme }}"
```

---

## ✅ Font Awesome Mapping

Dashuni supports *optional* FA icon mappings for dashboards like **Homer** that don't accept icon URLs.

⭐ Simply add this comment to the top of your template:

```gotemplate
{{/* requiresFA: true */}}
```

✅ Dashuni will automatically:
- Detect the header
- Load `./mappings/fa-mapping.json`
- Make the mapping available to your template as `.FAMap`

✅ Example template logic:

```gotemplate
{{ $faIcon := index $.FAMap .Title }}
{{ if $faIcon }}
icon: "{{ $faIcon }}"
{{ else }}
icon: "{{ .Icon }}"
{{ end }}
```

---

## ✅ ./mappings/fa-mapping.json Example

```json
{
  "Dockge": "fas fa-server",
  "FreshRSS": "fas fa-rss",
  "VaultWarden": "fas fa-lock"
}
```

---
## ⚠️ Notes for alternate icon sets

When generating some configs (e.g. Mafl), Dashuni outputs `icon` entries like this:

```yaml
icon:
  name: "simple-icons:<servicename>"
  wrap: true
```

**Important:**
- Some dashboards (e.g. Mafl) use icon libraries such as **Simple Icons**, **Material Design Icons (mdi)**, and others.
- **Simple Icons** includes *only well-known brands*. Many self-hosted services (e.g., Dockge, FreshRSS, Hoarder) **do not have Simple Icons entries.**

As a result:
- If you use `simple-icons:<servicename>` for unsupported services, the icon **will not render** and you will see a fallback text badge.
- You can solve this by:
  - Replacing `simple-icons:<servicename>` with a known icon from another library (e.g., `mdi:docker`, `mdi:rss`).
  - Manually uploading custom icons in the dashboard UI and adjusting the config to reference them.
  - Maintaining an **icon mapping file** and customizing the output template to use it.

If you want your dashboards to display icons consistently, it is recommended to:
- Test your generated config in MAFL.
- Adjust icon names to match supported icon sets.
- Consider using an icon mapping strategy to override defaults.

Refer to the MAFL documentation for a list of supported icons.

---

## ✅ Writing Templates

- Templates can embed mapping logic, target mapping, etc.  
- Example target mapping in Dashy:

```gotemplate
{{ $targetMap := dict "_blank" "newtab" "_self" "sametab" "_top" "top" }}
target: "{{ index $targetMap .Target }}"
```

---

## ✅ Example Commands

✅ For Dashy:

```bash
./dashuni convert --input sample.json --mapping mappings/dashy.tmpl --output dashy.yml
```

✅ For Homer (auto-detects FA mapping):

```bash
./dashuni convert --input sample.json --mapping mappings/homer.tmpl --output homer.yml
```

---

## ✅ Contributing

⭐ PRs welcome  
⭐ Add new dashboard templates  
⭐ Help refine the schema  

---

## 📜 License

MIT