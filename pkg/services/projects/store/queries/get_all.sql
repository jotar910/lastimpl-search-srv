SELECT p.id,
    p.name,
    p.description,
    cf.id,
    cf.name,
    cf.content,
    ptag.tag
FROM (
        SELECT id,
            name,
            description
        FROM projects
        WHERE name ~ $1
        ORDER BY name
        LIMIT $2
        OFFSET $3
    ) AS p
    INNER JOIN code_files AS cf ON p.id = cf.project_id
    INNER JOIN projects_tags AS ptag ON p.id = ptag.project_id;