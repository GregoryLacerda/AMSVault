# Guia de Implementação das Melhorias do AMSVault

Este documento fornece um guia passo a passo para implementar as melhorias listadas no arquivo `improvements.todo`.

## Arquitetura e Design

### 1. Implementar tratamento de erros centralizado
1. Crie um pacote `errors` dentro de `pkg`
2. Defina tipos de erro personalizados para diferentes categorias (DB, API, Business Logic)
3. Implemente funções utilitárias para criar erros com contexto
4. Substitua erros genéricos por erros personalizados em todo o projeto
5. Implemente middleware de tratamento de erro para a API

### 2. Adicionar middleware de autenticação e autorização
1. Crie um pacote `middleware/auth` para autenticação e autorização
2. Implemente um middleware para validar JWT tokens
3. Adicione verificação de permissões baseada em roles
4. Aplique o middleware nas rotas que requerem autenticação
5. Implemente tratamento para tokens expirados

### 3. Substituir uso de `panic` por tratamento adequado de erros
1. Identifique todos os `panic` no código
2. Substitua por retorno de erro com contexto apropriado
3. Implemente tratamento de erros nos chamadores
4. Adicione logging para erros críticos
5. Garanta que a aplicação continue funcionando após erros não-fatais

### 4. Implementar interfaces para dependências
1. Defina interfaces para todos os serviços e repositórios
2. Refatore implementações concretas para satisfazer as interfaces
3. Use injeção de dependência através das interfaces
4. Atualize os testes para usar mocks das interfaces
5. Documente o propósito de cada interface

### 5. Refatorar organização de pastas (Clean Architecture)
1. Separe claramente as camadas: apresentação, casos de uso, domínio, infraestrutura
2. Mova entidades para a camada de domínio
3. Reorganize services como casos de uso
4. Extraia interfaces de repositório para a camada de domínio
5. Garanta que dependências apontem para dentro (em direção ao domínio)

### 6. Implementar versionamento da API
1. Adicione prefixo de versão nas rotas (ex: `/api/v1/`)
2. Organize controladores por versão
3. Implemente sistema para gerenciar diferentes versões simultaneamente
4. Documente mudanças de API entre versões
5. Adicione cabeçalhos de versão nas respostas

### 7. Separar entidades de domínio dos DTOs
1. Crie modelos específicos para requisições e respostas (DTOs)
2. Implemente mapeadores/transformadores entre DTOs e entidades de domínio
3. Valide DTOs na camada de apresentação
4. Mantenha entidades de domínio focadas em regras de negócio
5. Remova campos específicos de apresentação das entidades de domínio

### 8. Implementar graceful shutdown
1. Adicione tratamento para sinais do sistema operacional (SIGTERM, SIGINT)
2. Implemente timeout para conexões ativas se encerrarem
3. Feche corretamente conexões com banco de dados
4. Notifique clientes conectados sobre o shutdown
5. Adicione logs para o processo de shutdown

## Banco de Dados

### 1. Implementar migrações de banco de dados
1. Instale uma ferramenta de migração (golang-migrate, goose, etc.)
2. Crie scripts de migração para o esquema atual
3. Adicione migrações para cada alteração de esquema
4. Implemente rotina para executar migrações na inicialização
5. Documente o processo de criação de novas migrações

### 2. Adicionar índices para campos frequentemente consultados
1. Identifique campos frequentemente usados em WHERE, JOIN e ORDER BY
2. Adicione índices para `name` e `mal_id` na tabela `stories`
3. Crie índices compostos para consultas comuns
4. Documente índices no esquema do banco de dados
5. Meça o impacto de performance antes e depois

### 3. Implementar transações de banco de dados
1. Identifique operações que modificam múltiplas tabelas
2. Refatore o código para usar transações SQL 
3. Implemente rollback em caso de erro
4. Adicione commit apenas quando todas as operações forem bem-sucedidas
5. Teste cenários de falha para garantir integridade

### 4. Adicionar soft delete consistente
1. Adicione campo `deleted_at` em todas as tabelas relevantes
2. Modifique operações DELETE para apenas marcar registros como deletados
3. Atualize todas as consultas para filtrar registros deletados
4. Implemente funcionalidade para exclusão permanente quando necessário
5. Documente a estratégia de soft delete

## Segurança

### 1. Implementar rate limiting
1. Adicione middleware para monitorar frequência de requisições
2. Implemente algoritmo de rate limiting (token bucket, leaky bucket)
3. Configure limites diferentes por tipo de endpoint
4. Adicione respostas apropriadas para quando limite for excedido
5. Implemente sistema para whitelisting de IPs ou clientes confiáveis

### 2. Adicionar validação de entrada
1. Implemente validação de todos os dados de entrada na API
2. Adicione sanitização para prevenir injeção SQL e XSS
3. Valide tipos, formatos e limites dos campos
4. Crie respostas de erro detalhadas para validação falha
5. Teste com dados maliciosos para validar segurança

### 3. Implementar proteção CSRF
1. Adicione geração de tokens CSRF para formulários e operações sensíveis
2. Implemente verificação de tokens para requisições de modificação
3. Adicione cabeçalhos de proteção contra CSRF
4. Configure política SameSite para cookies
5. Teste a proteção com ferramentas de segurança

### 4. Configurar headers de segurança
1. Adicione Content-Security-Policy para restringir fontes de conteúdo
2. Configure X-XSS-Protection para browsers mais antigos
3. Adicione X-Content-Type-Options: nosniff
4. Configure Strict-Transport-Security para forçar HTTPS
5. Adicione X-Frame-Options para prevenir clickjacking

## Funcionalidades

### 1. Implementar busca avançada
1. Crie endpoints para busca avançada com múltiplos parâmetros
2. Implemente filtragem por categoria, status, gênero, etc.
3. Adicione ordenação e paginação
4. Implemente busca full-text para descrições
5. Crie interface para construção de consultas complexas

### 2. Adicionar sistema de recomendação
1. Colete dados de visualização e interação dos usuários
2. Implemente algoritmo simples de recomendação baseada em conteúdo
3. Sugira títulos baseados nos gêneros favoritos do usuário
4. Adicione recomendações baseadas em popularidade
5. Implemente feedback para melhorar recomendações

### 3. Implementar notificações
1. Crie sistema de eventos para novos episódios/capítulos
2. Implemente fila de notificações
3. Adicione opção de preferências de notificação para usuários
4. Implemente múltiplos canais (email, push, in-app)
5. Crie templates para diferentes tipos de notificações

## Testes

### 1. Adicionar testes unitários
1. Configure ambiente de teste com ferramentas apropriadas
2. Comece pelos serviços de domínio mais críticos
3. Use mocks para dependências externas
4. Implemente testes para casos de sucesso e falha
5. Automatize execução de testes unitários no processo de build

### 2. Implementar testes de integração
1. Configure ambiente isolado para testes de integração
2. Implemente testes para operações de banco de dados
3. Teste fluxos completos envolvendo múltiplos componentes
4. Use containers para simular serviços externos
5. Adicione limpeza de dados após cada teste

### 3. Criar testes end-to-end
1. Configure ferramenta para testes API (Postman, Newman, etc.)
2. Crie coleção de testes para principais fluxos da aplicação
3. Simule ações de usuário através da API
4. Verifique respostas completas incluindo headers e status codes
5. Automatize execução dos testes end-to-end no CI

## Documentação

### 1. Criar documentação da API
1. Instale e configure Swagger/OpenAPI
2. Adicione anotações nos controladores para gerar documentação
3. Descreva parâmetros, respostas e códigos de erro
4. Inclua exemplos para cada endpoint
5. Publique documentação em ambiente de desenvolvimento/teste

### 2. Documentar estrutura do banco de dados
1. Crie diagrama ER do banco de dados
2. Documente propósito de cada tabela
3. Liste índices e restrições
4. Explique relações entre tabelas
5. Mantenha a documentação atualizada com cada migração

### 3. Adicionar README completo
1. Descreva visão geral do projeto
2. Documente requisitos de sistema
3. Inclua instruções detalhadas de instalação e configuração
4. Adicione exemplos de uso
5. Liste tecnologias utilizadas e suas versões

## Qualidade do Código

### 1. Padronizar mensagens de erro
1. Mova todas as strings de erro para constantes
2. Organize constantes por domínio ou componente
3. Padronize formato das mensagens
4. Inclua códigos de erro para facilitar rastreamento
5. Documente códigos de erro importantes

### 2. Aplicar linting e formatação
1. Configure ferramentas de linting (golangci-lint)
2. Defina regras de formatação e estilo
3. Integre linting no processo de build
4. Corrija problemas identificados pelo linter
5. Adicione verificação de linting no CI

### 3. Remover código duplicado
1. Identifique padrões repetidos no código
2. Extraia funções utilitárias reutilizáveis
3. Implemente abstrações para comportamentos comuns
4. Utilize composição e herança quando apropriado
5. Refatore para aplicar o princípio DRY

## DevOps

### 1. Configurar CI/CD
1. Configure GitHub Actions ou outra ferramenta de CI/CD
2. Automatize build e testes
3. Implemente verificações de qualidade de código
4. Configure deployment automatizado para ambientes de teste
5. Implemente gates de qualidade para produção

### 2. Containerizar a aplicação
1. Crie Dockerfile otimizado para a aplicação
2. Configure multi-stage build para reduzir tamanho da imagem
3. Utilize docker-compose para desenvolvimento local
4. Implemente healthchecks no container
5. Configure volumes para dados persistentes

### 3. Implementar logs estruturados
1. Adote biblioteca de logging estruturado (zap, logrus, etc.)
2. Padronize campos de log (request ID, user ID, etc.)
3. Configure níveis de log apropriados
4. Implemente rotação de logs
5. Configure agregação de logs para ambiente de produção

## Performance

### 1. Implementar paginação
1. Adicione parâmetros de paginação (limit, offset/cursor)
2. Implemente paginação em todos os endpoints de listagem
3. Otimize consultas SQL para paginação eficiente
4. Retorne metadados de paginação nas respostas
5. Implemente controle de tamanho máximo de página

### 2. Adicionar cache
1. Identifique dados frequentemente acessados e pouco alterados
2. Implemente cache em memória para consultas frequentes
3. Configure tempos de expiração adequados
4. Implemente invalidação de cache após modificações
5. Monitore hit ratio do cache para ajustes

### 3. Otimizar consultas ao banco
1. Analise logs de consultas lentas
2. Otimize joins e subconsultas
3. Utilize EXPLAIN para analisar planos de execução
4. Ajuste índices conforme necessário
5. Considere views materializadas para consultas complexas

## Específicos do Projeto

### 1. Adicionar mais fontes de dados
1. Pesquise APIs disponíveis para séries (TMDB, TVMaze)
2. Implemente clientes para as novas APIs
3. Padronize modelo de dados para unificar diferentes fontes
4. Adicione priorização de fontes de dados
5. Implemente fallback entre fontes em caso de falha

### 2. Implementar sincronização de status
1. Adicione função para exportar status para serviços externos
2. Implemente sincronização bidirecional
3. Gerencie conflitos de dados
4. Adicione opção para usuários controlarem sincronização
5. Implemente histórico de sincronização

### 3. Adicionar listas personalizadas
1. Crie modelo para listas personalizadas
2. Implemente CRUD de listas
3. Adicione funcionalidade para adicionar/remover itens
4. Implemente ordenação personalizada de itens
5. Adicione compartilhamento de listas entre usuários
