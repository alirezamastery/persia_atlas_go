package rawsql

var SqlVariant = `
SELECT
    v.id
  , v.dkpc
  , v.product_id
  , v.stop_loss
  , v.is_active
  , v.has_competition
  , p.id           AS "Product__id"
  , p.dkp          AS "Product__dkp"
  , p.dkp          AS "Product__dkp"
  , p.title        AS "Product__title"
  , p.is_active    AS "Product__is_active"
  , pt.id          AS "Product__Type__id"
  , pt.title       AS "Product__Type__title"
  , vst.id         AS "Product__Type__SelectorType__id"
  , vst.title      AS "Product__Type__SelectorType__title"
  , vs.id          AS "Selector__id"
  , vs.digikala_id AS "Selector__digikala_id"
  , vs.value       AS "Selector__value"
  , vs.extra_info  AS "Selector__extra_info"
  , ap.id          AS ActualProduct__id
  , ap.title       AS ActualProduct__title
  , ap.price_step  AS ActualProduct__price_step
  , b.id           AS ActualProduct__Brand__id
  , b.title        AS ActualProduct__Brand__title
FROM
    variants                          AS v
    INNER JOIN products               AS p
               ON v.product_id = p.id
    INNER JOIN product_types          AS pt
               ON p.type_id = pt.id
    INNER JOIN variant_selectors      AS vs
               ON v.selector_id = vs.id
    INNER JOIN variant_selector_types AS vst
               ON vs.selector_type_id = vst.id
    INNER JOIN actual_products        AS ap
               ON v.actual_product_id = ap.id
    INNER JOIN brands                 AS b
               ON ap.brand_id = b.id
WHERE
    v.id = ?
ORDER BY
    v.id
LIMIT 1;
`
