[简体中文](README.zh_CN.md)

## Github starred repo backup
This application can backup your starred github repos save as YAML file.  
YAML file structure like this:
```YAML
- repo name: DRQN_Stock_Trading
  url: https://github.com/conditionWang/DRQN_Stock_Trading
  description: 'This is the code implementation of the paper "Financial Trading as
    a Game: A Deep Reinforcement Learning Approach".'
  author: conditionWang
  author's github: https://github.com/conditionWang
  language: Python
  tags: drqn, lstm, reinforcement-learning, stock-trading
- repo name: MES40
  url: https://github.com/KH40-khoast40/MES40
  description: MMD Extended Shader (MES40)
  author: KH40-khoast40
  author's github: https://github.com/KH40-khoast40
  language: HLSL
  tags: diva, fx, hlsl, materials, mikumikudance, mikumikueffect, mmd, mme, project,
    shader, shadow
```

## Dependcies
This app use gopkg.in/yaml.v2 to marshal serializes the repo struct provided into a YAML file.

## How to use
### application parameters:
```bash
-u [your github username]
-p [the directory you want to save to. (optional)]
-m [the maxinum number of your starred repos]
```

### Usage example:
```bash
go run . -u=nikusaikou -p=stars -m=140
```
