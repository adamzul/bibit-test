SELECT a.ID, a.UserName, b.UserName AS Parent FROM user a
LEFT JOIN user b ON a.Parent = b.ID