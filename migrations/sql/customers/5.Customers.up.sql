CREATE TABLE IF NOT EXISTS  public.customers (
  id            UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name          TEXT NOT NULL,
  lastname      TEXT NOT NULL,  
  phone         VARCHAR(32),
  email         VARCHAR(128),
  address       TEXT,
  tags          JSONB,
  images        UUID REFERENCES public.images (id), 
  currency_id   UUID REFERENCES public.currency (id),
  created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP WITHOUT TIME ZONE,
  is_deleted    BOOLEAN NOT NULL DEFAULT FALSE
);
