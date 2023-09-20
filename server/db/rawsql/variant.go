package rawsql

var SqlVariant = `
SELECT
    v.id               AS "Id"
  , v.dkpc             AS "Dkpc"
  , v.product_id       AS "Product_id"
  , v.stop_loss        AS "Stop_loss"
  , v.price_min        AS "Price_min"
  , v.is_active        AS "Is_active"
  , v.has_competition  AS "Has_competition"
  , p.id               AS "Product__id"
  , p.dkp              AS "Product__dkp"
  , p.title            AS "Product__title"
  , p.is_active        AS "Product__is_active"
  , pt.id              AS "Product__Type__id"
  , pt.title           AS "Product__Type__title"
  , vst.id             AS "SelectorType__id"
  , vst.title          AS "SelectorType__title"
  , vs.id              AS "Selector__id"
  , vs.digikala_id     AS "Selector__digikala_id"
  , vs.value           AS "Selector__value"
  , vs.extra_info      AS "Selector__extra_info"
  , ap.id              AS "ActualProduct__id"
  , ap.title           AS "ActualProduct__title"
  , ap.price_step      AS "ActualProduct__price_step"
  , b.id               AS "ActualProduct__Brand__id"
  , b.title            AS "ActualProduct__Brand__title"
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

var SqlVariant2 = `
SELECT
    v.id               AS "Id"
  , v.dkpc             AS "Dkpc"
  , v.product_id       AS "ProductId"
  , v.stop_loss        AS "StopLoss"
  , v.price_min        AS "PriceMin"
  , v.is_active        AS "IsActive"
  , v.has_competition  AS "HasCompetition"
  , p.id               AS "ProductId"
  , p.dkp              AS "ProductDkp"
  , p.title            AS "ProductTitle"
  , p.is_active        AS "ProductIsActive"
  , pt.id              AS "ProductTypeId"
  , pt.title           AS "ProductTypeTitle"
  , vst.id             AS "SelectorTypeId"
  , vst.title          AS "SelectorTypeTitle"
  , vs.id              AS "SelectorId"
  , vs.digikala_id     AS "SelectorDigikalaId"
  , vs.value           AS "SelectorValue"
  , vs.extra_info      AS "SelectorExtraInfo"
  , ap.id              AS "ActualProductId"
  , ap.title           AS "ActualProductTitle"
  , ap.price_step      AS "ActualProductPriceStep"
  , b.id               AS "BrandId"
  , b.title            AS "BrandTitle"
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

var SqlVariantNoName = `
SELECT
    v.id             
  , v.dkpc           
  , v.product_id     
  , v.stop_loss      
  , v.price_min      
  , v.is_active      
  , v.has_competition
  , p.id             
  , p.dkp            
  , p.title          
  , p.is_active      
  , pt.id            
  , pt.title         
  , vst.id           
  , vst.title        
  , vs.id            
  , vs.digikala_id   
  , vs.value         
  , vs.extra_info    
  , ap.id            
  , ap.title         
  , ap.price_step    
  , b.id             
  , b.title          
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
