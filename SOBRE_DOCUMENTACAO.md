# Sobre a Documentação do Projeto AMSVault

## Arquivos de Documentação

Este projeto possui três arquivos principais de documentação projetados para auxiliar tanto desenvolvedores humanos quanto agentes de IA a entenderem e trabalharem com o código:

### 1. [`agents.md`](agents.md) - Guia Geral para Agentes de IA

**Propósito**: Visão geral completa da arquitetura e estrutura do projeto

**Conteúdo**:
- Arquitetura em camadas detalhada
- Estrutura completa de diretórios
- Stack tecnológico e dependências
- Fluxo de inicialização da aplicação
- Padrões de código para cada camada
- Modelos de dados (User, Story, Bookmarks)
- Sistema de autenticação JWT
- Rotas da API
- Configuração de ambiente
- Convenções de nomenclatura
- Comandos úteis
- Guia para adicionar novas features

**Quando usar**: Para entender a estrutura geral do projeto ou quando precisar adicionar uma nova entidade/feature completa.

---

### 2. [`claude.md`](claude.md) - Contexto Detalhado para Claude (e outros LLMs)

**Propósito**: Contexto específico com foco em padrões de implementação e boas práticas

**Conteúdo**:
- Filosofia do projeto (Clean Architecture, DDD, etc.)
- Fluxo de dados através das camadas
- Regras de validação por camada
- Padrões de implementação com exemplos práticos
- Diferenças entre MySQL e MongoDB
- Sistema de autenticação detalhado
- ViewModels e DTOs
- Constantes e erros customizados
- Comandos de desenvolvimento
- Debugging e erros comuns
- Modificações comuns (com exemplos de código)
- FAQs
- Checklist de modificação
- Filosofia de código

**Quando usar**: Para entender como implementar algo seguindo os padrões do projeto ou quando tiver dúvidas sobre a melhor abordagem.

---

### 3. [`implementation_guide.md`](implementation_guide.md) - Guia de Implementação de Melhorias

**Propósito**: Roadmap prático de melhorias e features a implementar

**Conteúdo**:
- Status de cada feature (✅ Implementado, ⚠️ Parcial, 🔴 Não implementado)
- Instruções passo a passo para implementações
- Código completo e pronto para usar
- Priorização de tarefas
- Checklist de progresso
- Exemplos práticos de:
  - Sistema de migrações de BD
  - Rate limiting
  - Validação robusta
  - Testes unitários e de integração
  - CI/CD com GitHub Actions
  - Logs estruturados
  - Swagger/OpenAPI
  - E muito mais...

**Quando usar**: Quando quiser implementar uma melhoria específica ou quando não souber o que fazer a seguir.

---

## Fluxo de Trabalho Recomendado

### Para Novos Desenvolvedores/IAs:

1. **Primeira vez no projeto?**
   - Leia [`agents.md`](agents.md) para entender a estrutura geral
   - Veja os diagramas de arquitetura
   - Entenda o fluxo de inicialização

2. **Vai implementar algo novo?**
   - Consulte [`claude.md`](claude.md) para ver padrões e exemplos
   - Verifique seção de "Modificações Comuns"
   - Siga o checklist de modificação

3. **Procurando o que implementar?**
   - Abra [`implementation_guide.md`](implementation_guide.md)
   - Veja a seção de priorização
   - Escolha uma tarefa e siga os passos

4. **Durante o desenvolvimento:**
   - Use [`claude.md`](claude.md) como referência de padrões
   - Consulte [`agents.md`](agents.md) para estrutura de pastas
   - Siga os exemplos em [`implementation_guide.md`](implementation_guide.md)

---

## Características Importantes dos Arquivos

### Formato Otimizado para IAs

Todos os arquivos foram escritos com foco em:
- **Clareza**: Linguagem direta e sem ambiguidades
- **Exemplos Práticos**: Código real, não pseudocódigo
- **Contexto**: Explica o "porquê" além do "como"
- **Estrutura**: Hierarquia clara com títulos e subtítulos
- **Status Explícito**: Indica o que já existe vs. o que precisa ser feito

### Código Pronto para Uso

Os exemplos de código em [`implementation_guide.md`](implementation_guide.md) são:
- ✅ Completos e funcionais
- ✅ Seguem os padrões do projeto
- ✅ Incluem imports necessários
- ✅ Testados e validados
- ✅ Comentados quando necessário

### Atualização Contínua

Estes arquivos devem ser atualizados quando:
- Nova feature for implementada
- Arquitetura mudar
- Novos padrões forem adotados
- Tecnologias forem adicionadas/removidas

---

## Integração com Outras Documentações

### [`API_DOCUMENTATION.md`](API_DOCUMENTATION.md)
Documentação completa da API REST com todos os endpoints, exemplos de requisições e respostas.

### `README.md`
(Se existir) Instruções de instalação e setup inicial do projeto.

### Código Fonte
O código em si é a documentação definitiva. Use os arquivos md como guia, mas sempre valide no código real.

---

## Para Mantenedores

### Ao Adicionar Nova Feature

1. Atualize [`agents.md`](agents.md):
   - Adicione na estrutura de diretórios se criar novas pastas
   - Atualize seção de rotas se adicionar endpoints
   - Documente novos padrões se introduzir algum

2. Atualize [`claude.md`](claude.md):
   - Adicione exemplos na seção "Modificações Comuns"
   - Atualize FAQs se houver dúvidas frequentes
   - Documente novos padrões de implementação

3. Atualize [`implementation_guide.md`](implementation_guide.md):
   - Marque como ✅ se era uma tarefa pendente
   - Remova da lista de tarefas
   - Atualize checklist

### Ao Mudar Arquitetura

1. Revise TODOS os três arquivos
2. Atualize diagramas se houver
3. Corrija exemplos de código obsoletos
4. Atualize seções de fluxo de dados

---

## Dicas para IAs

### Ao Receber Uma Solicitação:

1. **Identifique o tipo de tarefa**:
   - Entender o projeto? → [`agents.md`](agents.md)
   - Implementar algo específico? → [`implementation_guide.md`](implementation_guide.md)
   - Dúvida sobre padrão? → [`claude.md`](claude.md)

2. **Verifique o status atual**:
   - Leia a seção relevante em [`implementation_guide.md`](implementation_guide.md)
   - Confirme se já existe ou não

3. **Siga os padrões**:
   - Use exemplos de [`claude.md`](claude.md) como referência
   - Mantenha consistência com código existente

4. **Valide contra a documentação**:
   - Verifique se a implementação segue os padrões
   - Confirme que atualiza onde necessário

---

## Estrutura de Prioridades

### Alta Prioridade (Fazer Primeiro)
Features críticas para segurança, estabilidade e funcionalidade básica.

### Média Prioridade
Melhorias importantes mas não críticas.

### Baixa Prioridade
Nice-to-have, pode esperar.

Veja detalhes completos em [`implementation_guide.md`](implementation_guide.md#priorização-sugerida).

---

## Contribuindo

Ao contribuir com o projeto, por favor:

1. ✅ Leia a documentação relevante antes
2. ✅ Siga os padrões estabelecidos
3. ✅ Adicione testes quando possível
4. ✅ Atualize a documentação se necessário
5. ✅ Faça commits descritivos

---

## Suporte

Para questões sobre a documentação:
- Abra uma issue no repositório
- Sugira melhorias via pull request
- Reporte erros ou inconsistências

---

**Última Atualização**: Fevereiro 2026
**Versão da Documentação**: 1.0
