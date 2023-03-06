[简体中文](README.zh_CN.md)

## backup GitHub starred repos
This application can backup your starred GitHub repos, saving them as a YAML file.  
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
This app uses gopkg.in/yaml.v2 to serialize the repo struct into a YAML file.

## How to use
### application parameters:
```bash
-u [your github username]
-p [the directory you want to save to. (optional)]
-m [the maxinum number of your starred repos]
-g [use goroutine or not true/false default: true(optional. if true the result is unordered))]
```

### Usage example:
```bash
go run . -u=nikusaikou -p=stars -m=140 -g=false
```
