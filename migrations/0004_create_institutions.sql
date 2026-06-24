CREATE TABLE IF NOT EXISTS institutions (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name TEXT NOT NULL,
  abbreviation TEXT UNIQUE NOT NULL,
  city TEXT NOT NULL,
  address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO institutions (name, abbreviation, city, address) VALUES
  ('Universidade Federal do Rio de Janeiro', 'UFRJ', 'Rio de Janeiro', 'Av. Pedro Calmon, 550, Cidade Universitária, Rio de Janeiro - RJ, 21941-901'),
  ('Universidade Federal Fluminense', 'UFF', 'Niterói', 'Rua Miguel de Frias, 9, Icaraí, Niterói - RJ, 24220-900'),
  ('Universidade Federal Rural do Rio de Janeiro', 'UFRRJ', 'Seropédica', 'Rod. BR-465, Km 7, Zona Rural, Seropédica - RJ, 23890-000'),
  ('Universidade Federal do Estado do Rio de Janeiro', 'UNIRIO', 'Rio de Janeiro', 'Av. Pasteur, 296, Urca, Rio de Janeiro - RJ, 22290-240'),
  ('Instituto Federal do Rio de Janeiro', 'IFRJ', 'Rio de Janeiro', 'Rua Pereira de Almeida, 88, Praça da Bandeira, Rio de Janeiro - RJ, 20260-100'),
  ('Instituto Federal Fluminense', 'IFF', 'Campos dos Goytacazes', 'Rua Dr. Siqueira, 273, Parque Dom Bosco, Campos dos Goytacazes - RJ, 28030-130'),
  ('Centro Federal de Educação Tecnológica Celso Suckow da Fonseca', 'CEFET/RJ', 'Rio de Janeiro', 'Av. Maracanã, 229, Maracanã, Rio de Janeiro - RJ, 20271-110'),
  ('Colégio Pedro II', 'CPII', 'Rio de Janeiro', 'Campo de São Cristóvão, 177, São Cristóvão, Rio de Janeiro - RJ, 20921-440'),
  ('Fundação Oswaldo Cruz', 'FIOCRUZ', 'Rio de Janeiro', 'Av. Brasil, 4365, Manguinhos, Rio de Janeiro - RJ, 21040-900'),
  ('Instituto de Matemática Pura e Aplicada', 'IMPA', 'Rio de Janeiro', 'Estrada Dona Castorina, 110, Jardim Botânico, Rio de Janeiro - RJ, 22460-320'),
  ('Centro Brasileiro de Pesquisas Físicas', 'CBPF', 'Rio de Janeiro', 'Rua Dr. Xavier Sigaud, 150, Urca, Rio de Janeiro - RJ, 22290-180'),
  ('Escola Nacional de Ciências Estatísticas', 'ENCE', 'Rio de Janeiro', 'Rua André Cavalcanti, 106, Santa Teresa, Rio de Janeiro - RJ, 20231-050'),
  ('Universidade do Estado do Rio de Janeiro', 'UERJ', 'Rio de Janeiro', 'Rua São Francisco Xavier, 524, Maracanã, Rio de Janeiro - RJ, 20550-013'),
  ('Universidade Estadual do Norte Fluminense Darcy Ribeiro', 'UENF', 'Campos dos Goytacazes', 'Av. Alberto Lamego, 2000, Parque Califórnia, Campos dos Goytacazes - RJ, 28013-602'),
  ('Fundação de Apoio à Escola Técnica', 'FAETEC', 'Rio de Janeiro', 'Rua Clarimundo de Melo, 847, Quintino Bocaiúva, Rio de Janeiro - RJ, 21311-280');
