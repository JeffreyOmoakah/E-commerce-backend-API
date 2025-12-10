-- name: ListProducts :many 
SELECT 
 *
FROM 
    Products;

-- name: FindProductbyID :one
SELECT 
 *
FROM 
    Products
WHERE 
    id = $1;