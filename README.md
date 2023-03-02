[简体中文](README.zh_CN.md)

## Github starred repo backup
This application can backup your github repos save as YAML file.  
like this:
```YAML
- repo name: DRQN_Stock_Trading
  author: conditionWang
  author's github: https://github.com/conditionWang
  url: https://github.com/conditionWang/DRQN_Stock_Trading
  description: ""
  language: Python
  tags: drqn, lstm, reinforcement-learning, stock-trading
- repo name: podcast
  author: bumingbaipod
  author's github: https://github.com/bumingbaipod
  url: https://github.com/bumingbaipod/podcast
  description: ""
  language: ""
  tags: ""
```

## Dependcies
This app use gopkg.in/yaml.v2 to marshal serializes the repo struct provided into a YAML file.

## How to use
### application parameters:
```bash
-u [your github username]
-p [the directory you want to save to]
-m [the maxinum number of your starred repos]
```

### Usage example:
```bash
go run . -u=nikusaikou -m=141
```
