-- Initial Values

INSERT INTO pet_status (pet_status_id, pet_status_name)
VALUES
  (10, 'available'),
  (20, 'pending'),
  (30, 'sold')
ON CONFLICT (pet_status_id)
DO NOTHING;

INSERT INTO order_status (order_status_id, order_status_name)
VALUES
(10, 'placed'),
(20, 'approved'),
(30, 'delivered')
ON CONFLICT (pet_status_id)
DO NOTHING;
