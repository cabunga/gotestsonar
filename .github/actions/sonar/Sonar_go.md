# action-sonar-go

Action creado para enviar reporte de cobertura de los test y hallazgos en el codigo en sonarcloud


## Inputs

> Requeridos si usás el template sin los jobs previos.

| Input | Descripción |
|---|---|
| `organization` | Nombre de la organizacion (ej: `claro-video`) |
| `projectKey` | Nombre del Proyecto (ej: `Device`) |
| `sonarToken` | Secreto de SonarCloud (ej: `XXXXX`) |


---





## Secrets requeridos

> El workflow usa `secrets: inherit` — los secrets se heredan del workflow que lo llama.

| Secret | Descripción |
|---|---|
| `SONAR_TOKEN` | Token generado a nivel de SonarCloud |

---

## Variables requeridas

| Variable | Descripción |
|---|---|
| `projectKey` | Key del Proyecto en SonarCloud |
| `organization` | Nombre de la organizacion vinculada en SonarCloud |

---

## Cómo llamarlo

```yaml
steps:
      - name: Ejecutar Scan Sonar
        uses: uses: clarovideo-argentina/automation-workflow/.github/actions/sonar@main        
        with:
          projectKey: ${{ inputs.PROYECT }}
          organization: ${{ inputs.ORGANIZATION }}
          sonarToken: ${{ secrets.SONAR_TOKEN }}     

```
