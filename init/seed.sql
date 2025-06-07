-- Insert categories
INSERT INTO categories (name) VALUES
('Pain Relief'),
('Allergy'),
('Antibiotics');

-- Insert medicines
INSERT INTO medicienes (name, category_id) VALUES
('Paracetamol', 1),
('Ibuprofen', 1),
('Cetirizine', 2),
('Amoxicillin', 3);

-- Insert coupons
INSERT INTO coupons (
    coupoun_code, discount_type, discount_value,
    terms_and_conditions, usage_type, valid_from, valid_until, minimum_purchase_amount
) VALUES
('PAIN10', 'percentage', 10.00, 'Applicable to pain relief meds only.', 'multi_use', NOW(), NOW() + INTERVAL '30 days', 100.00),
('ALLERGY5', 'fixed_amount', 5.00, 'Use on any allergy medicine.', 'single_use', NOW(), NOW() + INTERVAL '15 days', 50.00),
('GENERIC20', 'percentage', 20.00, 'Valid on all categories above 200.', 'multi_use', NOW(), NOW() + INTERVAL '60 days', 200.00);

-- Map coupons to medicines
INSERT INTO coupon_medicines (coupon_id, medicine_id) VALUES
(1, 1), -- PAIN10 for Paracetamol
(1, 2), -- PAIN10 for Ibuprofen
(2, 3); -- ALLERGY5 for Cetirizine

-- Map coupons to categories
INSERT INTO coupon_categories (coupon_id, category_id) VALUES
(1, 1), -- PAIN10 -> Pain Relief
(2, 2), -- ALLERGY5 -> Allergy
(3, 1), -- GENERIC20 -> Pain Relief
(3, 2), -- GENERIC20 -> Allergy
(3, 3); -- GENERIC20 -> Antibiotics
