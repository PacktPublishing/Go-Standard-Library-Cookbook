CREATE OR REPLACE FUNCTION format_name
        (firstname Text,lastname Text,age INT) RETURNS 
        VARCHAR AS $$
        BEGIN
          RETURN trim(firstname) ||' '||trim(lastname) ||' ('||age||')';
        END;
        $$ LANGUAGE plpgsql;