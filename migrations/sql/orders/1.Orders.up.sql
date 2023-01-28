CREATE TABLE IF NOT EXISTS  public.orders (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID REFERENCES public.customer (id),
  bucket_goods_id UUID REFERENCES public.bucket_goods (id),  
  created_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      TIMESTAMP WITHOUT TIME ZONE,
  is_deleted      BOOLEAN NOT NULL DEFAULT FALSE
);

-- For many2many order and goods
CREATE TABLE IF NOT EXISTS public.bucket_goods (
  id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order          UUID REFERENCES public.orders (id), 
  good_id_1      UUID REFERENCES public.goods (id),
  good_id_2      UUID REFERENCES public.goods (id),
  good_id_3      UUID REFERENCES public.goods (id),
  good_id_4      UUID REFERENCES public.goods (id),
  good_id_5      UUID REFERENCES public.goods (id),
  next_bucket_id UUID REFERENCES public.bucket_goods
);

CREATE TABLE IF NOT EXISTS public.orders_history (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order       UUID REFERENCES public.orders (id), 
  created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  rating      INT DEFAULT 1,
  CONSTRAINT  order_rating CASE (rating >= 0 AND rating <= 5)
)
