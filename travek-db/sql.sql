SELECT 
    *
FROM 
    country
HAVING
    country.id IN(
        SELECT
            CASE
                WHEN relation.first_country = 23 THEN relation.second_country
                WHEN relation.second_country = 23 THEN relation.first_country
            END AS country_id
        FROM
            relation
        WHERE
            (relation.first_country = 23 OR relation.second_country = 23) AND
            (relation.free_entry = true)
        HAVING
            country_id IN(
                (SELECT
                    CASE
                        WHEN road.first_country = 23 THEN road.second_country
                        WHEN road.second_country = 23 THEN road.first_country
                    END AS country_id
                FROM
                    road
                WHERE
                    (road.first_country = 23 OR road.second_country = 23) AND
                    (road.transport = "plain") AND
                    (road.time_hours <= 8)
                )
            )
    )
;
    

-- поиск по дорогам 
SELECT country_id as c_id, CONCAT(GROUP_CONCAT(tag_id), ',') FROM country_tags GROUP BY c_id  HAVING (select CONCAT(GROUP_CONCAT(tag_id), ',') from country_tags where country_id = c_id) LIKE "%1_%" ESCAPE ',';
SELECT country_id as c_id FROM country_tags GROUP BY c_id HAVING 1 IN((select tag_id from country_tags where country_id = c_id)) AND 8 IN((select tag_id from country_tags where country_id = c_id))
(SELECT
    CASE
        WHEN road.first_country = 23 THEN road.second_country
        WHEN road.second_country = 23 THEN road.first_country
    END AS country_id
FROM
    road
WHERE
    (road.first_country = 23 OR road.second_country = 23) AND
    (road.transport IN("plain", "train")) AND
    (road.time_hours <= 8)
)

-- поиск по тегам

-- tags = ("beautiful neature", "relax")

SELECT
    country_tags.country_id as country_id
FROM
    country_tags
WHERE
    country_tags.tag_id IN(
        (SELECT
            tags.id
        FROM
            tags
        WHERE
            tags.name IN("beautiful neature", "relax")
        )
    )
;

SELECT
	country.name
FROM
	country
WHERE
	country.id IN(
		(SELECT
			country_tags.country_id as country_id
		FROM
			country_tags
		WHERE
			country_tags.tag_id IN(
				2, 3, 4	
			)
		)
	)
;


-- поиск по отношениям

SELECT
    CASE
        WHEN relation.first_country = 23 THEN relation.second_country
        WHEN relation.second_country = 23 THEN relation.first_country
    END AS country_id
FROM
    relation
WHERE
    (relation.first_country = 23 OR relation.second_country = 23) AND
    (relation.free_entry = true) AND
    (relation.possibility_to_stay_forever = true) 
;

-- поиск по отношениям и дорогам

SELECT
    CASE
        WHEN relation.first_country = 23 THEN relation.second_country
        WHEN relation.second_country = 23 THEN relation.first_country
    END AS country_id
FROM
    relation
WHERE
    (relation.first_country = 23 OR relation.second_country = 23) AND
    (relation.free_entry = true)
HAVING
    country_id IN(
        (SELECT
            -- road.first_country,
            -- road.second_country
            CASE
                WHEN road.first_country = 23 THEN road.second_country
                WHEN road.second_country = 23 THEN road.first_country
            END AS country_id
        FROM
            road
        WHERE
            (road.first_country = 23 OR road.second_country = 23) AND
            (road.transport = "plain") AND
            (road.time_hours <= 8)
        )
    )
;