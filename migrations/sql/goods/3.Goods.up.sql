CREATE TABLE IF NOT EXISTS  public.goods (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name          TEXT NOT NULL,
  description   TEXT NOT NULL,
  rating        INT DEFAULT 0,
  specification JSONB,
  tags          JSONB,
  price_id      UUID REFERENCES public.prices (id),
  images        UUID REFERENCES public.images (id),    
  category_id   UUID REFERENCES public.category (id),
  created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP WITHOUT TIME ZONE,
  is_deleted    BOOLEAN NOT NULL DEFAULT FALSE,  
  CONSTRAINT    valid_rating CHECK (rating => 0 AND rating <=5 )
);

CREATE TABLE IF NOT EXISTS  public.category (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name          TEXT NOT NULL,
  tags          JSONB
);

CREATE TABLE IF NOT EXISTS  public.images (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name          TEXT,
  is_deleted    BOOLEAN NOT NULL DEFAULT FALSE,
  path          TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS  public.currency (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name          TEXT NOT NULL,
  symbol        VARCHAR(2) NOT NULL  
);

CREATE TABLE IF NOT EXISTS public.prices (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  price         BIGINT,  
  currency_id   UUID REFERENCES public.currency (id),
  created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT    positive_price CHECK (price => 0)
);

-- DATA --
INSERT INTO public.currency (name, symbol)
VALUES ('рубль', '₽');
INSERT INTO public.currency (name, symbol)
VALUES ('dollar', '$');

INSERT INTO public.category (name, tags)
VALUES ('3д принтеры', '{"aliases": ["3d", "additive", "print"]}'::jsonb );
INSERT INTO public.category (name)
VALUES ('роботы', '{"aliases": ["robot", "mech"]}'::jsonb);