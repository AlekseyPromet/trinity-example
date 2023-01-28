CREATE TABLE IF NOT EXISTS  public.orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id UUID REFERENCES public.customer (id),
  bucket_goods_id UUID REFERENCES public.bucket_goods (id),
  
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- For many2many order and goods
CREATE TABLE IF NOT EXISTS public.bucket_goods (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  good_id UUID REFERENCES public.goods (id),
);

CREATE TABLE IF NOT EXISTS public.orders_history (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order       UUID REFERENCES public.orders (id), 
  created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
  rating      INT DEFAULT 1,
  CONSTRAINT  order_rating CASE (rating >= 0 AND rating <= 5)
)
