-- Seed / mock data for local development and the Docker containers.
-- Institutions are already seeded by 0004_create_institutions.sql; this file
-- populates users, books, events and loans on top of them.
--
-- Every seeded user shares the password "estante123" (bcrypt hash below) so you
-- can log in with any of the emails. Foreign keys are resolved via subqueries
-- because the id columns are GENERATED ALWAYS AS IDENTITY.
--
-- Safe to re-run: existing rows are skipped via ON CONFLICT on the unique keys.

-- ---------------------------------------------------------------------------
-- Users
-- ---------------------------------------------------------------------------
INSERT INTO users (name, birth_date, email, password, address, document, cellphone, role, score, created_at, institution_id) VALUES
  ('Ana Beatriz Carvalho', '1988-03-12', 'ana.carvalho@estanteviva.com', '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua Voluntários da Pátria, 120, Botafogo, Rio de Janeiro - RJ', '11122233344', '21988887766', 'admin',   120, '2025-01-10 09:00:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFRJ')),
  ('Carlos Eduardo Lima',  '1979-07-25', 'carlos.lima@estanteviva.com',   '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Av. Visconde do Rio Branco, 880, Centro, Niterói - RJ',       '22233344455', '21988776655', 'teacher',  90,  '2025-01-18 14:30:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFF')),
  ('Mariana Souza',        '1985-11-02', 'mariana.souza@estanteviva.com', '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua Conde de Bonfim, 410, Tijuca, Rio de Janeiro - RJ',       '33344455566', '21987654321', 'teacher',  75,  '2025-02-03 10:15:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFRJ')),
  ('Pedro Henrique Alves', '2002-05-19', 'pedro.alves@estanteviva.com',   '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua Dias da Cruz, 55, Méier, Rio de Janeiro - RJ',            '44455566677', '21996543210', 'student',  40,  '2025-02-12 16:45:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFRJ')),
  ('Juliana Ribeiro',      '2001-09-30', 'juliana.ribeiro@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua São Francisco Xavier, 200, Maracanã, Rio de Janeiro - RJ','55566677788', '21991234567', 'student',  60,  '2025-02-20 11:00:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UERJ')),
  ('Rafael Mendes',        '2000-12-14', 'rafael.mendes@estanteviva.com', '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua Gavião Peixoto, 33, Icaraí, Niterói - RJ',                '66677788899', '21993456789', 'student',  25,  '2025-03-01 09:30:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFF')),
  ('Beatriz Fernandes',    '2003-01-08', 'beatriz.fernandes@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22','Av. Pasteur, 250, Urca, Rio de Janeiro - RJ',                 '77788899900', '21994567890', 'student',  55,  '2025-03-08 13:20:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UNIRIO')),
  ('Lucas Gabriel Santos', '2002-06-22', 'lucas.santos@estanteviva.com',  '$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Av. Maracanã, 300, Maracanã, Rio de Janeiro - RJ',            '88899900011', '21995678901', 'student',  10,  '2025-03-15 08:45:00-03', (SELECT id FROM institutions WHERE abbreviation = 'CEFET/RJ')),
  ('Fernanda Oliveira',    '2001-04-17', 'fernanda.oliveira@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22','Rua Pereira de Almeida, 90, Praça da Bandeira, Rio de Janeiro - RJ','99900011122','21996789012','student', 70, '2025-03-22 15:10:00-03', (SELECT id FROM institutions WHERE abbreviation = 'IFRJ')),
  ('Gustavo Pereira',      '1982-08-05', 'gustavo.pereira@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua Haddock Lobo, 120, Estácio, Rio de Janeiro - RJ',         '10020030040', '21997890123', 'teacher',  85,  '2025-04-02 10:00:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UERJ')),
  ('Camila Rodrigues',     '2003-10-11', 'camila.rodrigues@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22','Rod. BR-465, Km 7, Seropédica - RJ',                          '20030040050', '21998901234', 'student',  35,  '2025-04-10 12:00:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFRRJ')),
  ('Thiago Martins',       '2000-02-28', 'thiago.martins@estanteviva.com','$2a$10$yYygycs4SzG7oISmMtWJy.0YAeFqf5AeHgOosS4aL38h64Y1Npi22', 'Rua General Roca, 700, Tijuca, Rio de Janeiro - RJ',          '30040050060', '21999012345', 'student',  50,  '2025-04-18 17:30:00-03', (SELECT id FROM institutions WHERE abbreviation = 'UFRJ'))
ON CONFLICT (email) DO NOTHING;

-- ---------------------------------------------------------------------------
-- Books
--   Titles flagged 'lent' below have a matching active loan further down.
-- ---------------------------------------------------------------------------
INSERT INTO books (title, author, release_date, edition, status, created_at) VALUES
  ('Dom Casmurro',                          'Machado de Assis',          '1899-01-01', '1ª edição',  'lent',      '2025-01-05 09:00:00-03'),
  ('Memórias Póstumas de Brás Cubas',       'Machado de Assis',          '1881-01-01', '3ª edição',  'available', '2025-01-05 09:05:00-03'),
  ('O Cortiço',                             'Aluísio Azevedo',           '1890-01-01', '2ª edição',  'available', '2025-01-05 09:10:00-03'),
  ('Grande Sertão: Veredas',                'João Guimarães Rosa',       '1956-01-01', '1ª edição',  'lent',      '2025-01-05 09:15:00-03'),
  ('Vidas Secas',                           'Graciliano Ramos',          '1938-01-01', '5ª edição',  'available', '2025-01-05 09:20:00-03'),
  ('Capitães da Areia',                     'Jorge Amado',               '1937-01-01', '4ª edição',  'available', '2025-01-05 09:25:00-03'),
  ('A Hora da Estrela',                      'Clarice Lispector',         '1977-01-01', '2ª edição',  'lent',      '2025-01-05 09:30:00-03'),
  ('Iracema',                               'José de Alencar',           '1865-01-01', '3ª edição',  'available', '2025-01-05 09:35:00-03'),
  ('O Guarani',                             'José de Alencar',           '1857-01-01', '1ª edição',  'available', '2025-01-05 09:40:00-03'),
  ('Macunaíma',                             'Mário de Andrade',          '1928-01-01', '2ª edição',  'lent',      '2025-01-05 09:45:00-03'),
  ('Quincas Borba',                         'Machado de Assis',          '1891-01-01', '1ª edição',  'available', '2025-01-05 09:50:00-03'),
  ('Sagarana',                              'João Guimarães Rosa',       '1946-01-01', '2ª edição',  'available', '2025-01-05 09:55:00-03'),
  ('Laços de Família',                      'Clarice Lispector',         '1960-01-01', '1ª edição',  'lent',      '2025-01-05 10:00:00-03'),
  ('O Tempo e o Vento',                     'Erico Verissimo',           '1949-01-01', '3ª edição',  'available', '2025-01-05 10:05:00-03'),
  ('Triste Fim de Policarpo Quaresma',      'Lima Barreto',              '1915-01-01', '2ª edição',  'available', '2025-01-05 10:10:00-03'),
  ('Senhora',                               'José de Alencar',           '1875-01-01', '1ª edição',  'available', '2025-01-05 10:15:00-03'),
  ('Memórias de um Sargento de Milícias',   'Manuel Antônio de Almeida', '1854-01-01', '1ª edição',  'available', '2025-01-05 10:20:00-03'),
  ('O Ateneu',                              'Raul Pompeia',              '1888-01-01', '2ª edição',  'available', '2025-01-05 10:25:00-03'),
  ('Fogo Morto',                            'José Lins do Rego',         '1943-01-01', '1ª edição',  'available', '2025-01-05 10:30:00-03'),
  ('A Moreninha',                           'Joaquim Manuel de Macedo',  '1844-01-01', '6ª edição',  'available', '2025-01-05 10:35:00-03');

-- ---------------------------------------------------------------------------
-- Events
-- ---------------------------------------------------------------------------
INSERT INTO events (name, description, date, location, institution_id, created_at) VALUES
  ('Feira do Livro Universitário',   'Troca e doação de livros entre estudantes, com mesas de autógrafos e debates.', '2026-07-15 10:00:00-03', 'Centro de Convenções da UFRJ, Cidade Universitária', (SELECT id FROM institutions WHERE abbreviation = 'UFRJ'),   '2025-05-01 09:00:00-03'),
  ('Clube de Leitura: Clarice Lispector', 'Encontro mensal para discutir a obra de Clarice Lispector.',               '2026-07-03 18:30:00-03', 'Biblioteca Central, UFF, Niterói',                  (SELECT id FROM institutions WHERE abbreviation = 'UFF'),    '2025-05-05 11:00:00-03'),
  ('Sarau Literário de Inverno',     'Noite de poesia, leitura e música aberta à comunidade acadêmica.',             '2026-07-22 19:00:00-03', 'Auditório do CEFET/RJ, Maracanã',                   (SELECT id FROM institutions WHERE abbreviation = 'CEFET/RJ'),'2025-05-10 14:00:00-03'),
  ('Mutirão de Doação de Livros',    'Campanha de arrecadação de livros didáticos para a estante viva.',             '2026-06-28 09:00:00-03', 'Hall de entrada da UERJ, Maracanã',                 (SELECT id FROM institutions WHERE abbreviation = 'UERJ'),   '2025-05-15 10:30:00-03'),
  ('Roda de Conversa: Machado de Assis', 'Mediação sobre o realismo brasileiro e a obra de Machado de Assis.',       '2026-08-05 17:00:00-03', 'Sala 12, Instituto de Letras, UNIRIO',              (SELECT id FROM institutions WHERE abbreviation = 'UNIRIO'), '2025-05-20 16:00:00-03'),
  ('Oficina de Encadernação Artesanal', 'Aprenda a restaurar e encadernar livros usados.',                          '2026-08-12 14:00:00-03', 'Laboratório de Artes, IFRJ',                        (SELECT id FROM institutions WHERE abbreviation = 'IFRJ'),   '2025-05-25 13:00:00-03'),
  ('Encontro de Cordel e Literatura Popular', 'Apresentações de cordel e troca de folhetos.',                       '2026-09-02 15:30:00-03', 'Praça central da UFRRJ, Seropédica',                (SELECT id FROM institutions WHERE abbreviation = 'UFRRJ'),  '2025-06-01 09:45:00-03'),
  ('Lançamento da Estante Viva Digital', 'Apresentação do catálogo digital e como contribuir com acervos.',         '2026-09-18 11:00:00-03', 'Auditório da FIOCRUZ, Manguinhos',                  (SELECT id FROM institutions WHERE abbreviation = 'FIOCRUZ'),'2025-06-08 10:00:00-03');

-- ---------------------------------------------------------------------------
-- Loans
--   Active loans (returned_at NULL) keep their book in 'lent' status above.
--   Returned loans (returned_at set) point at books currently 'available'.
-- ---------------------------------------------------------------------------
INSERT INTO loans (user_id, book_id, return_date, returned_at) VALUES
  -- Active loans
  ((SELECT id FROM users WHERE email = 'pedro.alves@estanteviva.com'),     (SELECT id FROM books WHERE title = 'Dom Casmurro'),             '2026-07-05', NULL),
  ((SELECT id FROM users WHERE email = 'juliana.ribeiro@estanteviva.com'), (SELECT id FROM books WHERE title = 'Grande Sertão: Veredas'),   '2026-07-10', NULL),
  ((SELECT id FROM users WHERE email = 'rafael.mendes@estanteviva.com'),   (SELECT id FROM books WHERE title = 'A Hora da Estrela'),         '2026-07-02', NULL),
  ((SELECT id FROM users WHERE email = 'beatriz.fernandes@estanteviva.com'),(SELECT id FROM books WHERE title = 'Macunaíma'),               '2026-07-08', NULL),
  ((SELECT id FROM users WHERE email = 'lucas.santos@estanteviva.com'),    (SELECT id FROM books WHERE title = 'Laços de Família'),          '2026-06-30', NULL),
  -- Returned loans
  ((SELECT id FROM users WHERE email = 'pedro.alves@estanteviva.com'),     (SELECT id FROM books WHERE title = 'O Cortiço'),                '2026-05-20', '2026-05-18'),
  ((SELECT id FROM users WHERE email = 'juliana.ribeiro@estanteviva.com'), (SELECT id FROM books WHERE title = 'Vidas Secas'),              '2026-04-15', '2026-04-10'),
  ((SELECT id FROM users WHERE email = 'fernanda.oliveira@estanteviva.com'),(SELECT id FROM books WHERE title = 'Iracema'),                 '2026-06-05', '2026-06-01'),
  ((SELECT id FROM users WHERE email = 'thiago.martins@estanteviva.com'),  (SELECT id FROM books WHERE title = 'Capitães da Areia'),         '2026-03-20', '2026-03-19');
