CREATE TABLE work_experiences (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    workplace VARCHAR(255) NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  github_link VARCHAR(255),
  image_link VARCHAR(255),
  blog_link VARCHAR(255),
  type VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE skills (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  blog_link VARCHAR(255),
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);
