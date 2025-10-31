BEGIN;

-- 1) LABEL: assegna ciclicamente dove è NULL
UPDATE public.release r
SET label = CASE ((r.id - 1) % 4)
  WHEN 0 THEN 'metamorfosi'
  WHEN 1 THEN 'mood child'
  WHEN 2 THEN 'elite sound'
  ELSE      'hottrax'
END
WHERE r.label IS NULL;

-- 2) DATE: riempi i NULL con mesi consecutivi dopo la data massima esistente
WITH maxd AS (
  SELECT COALESCE(date_trunc('day', MAX("date")), '2024-01-01 00:00:00+00'::timestamptz) AS base
  FROM public.release
),
tofill AS (
  SELECT r.id, ROW_NUMBER() OVER (ORDER BY r.id) AS rn
  FROM public.release r
  WHERE r."date" IS NULL
)
UPDATE public.release r
SET "date" = maxd.base + (tofill.rn * INTERVAL '1 month')
FROM tofill, maxd
WHERE r.id = tofill.id;

COMMIT;
