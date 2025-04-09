-- migrations/2_add_software_table.sql
-- Add the 'software' table for storing software applications

-- Create software table
CREATE TABLE software (
    id VARCHAR(255) PRIMARY KEY,
    foreign_key VARCHAR(255),
    display_name VARCHAR(255) NOT NULL,
    description TEXT,
    software_type VARCHAR(50) NOT NULL, -- 'api', 'web', 'mobile', 'desktop', 'embedded', 'middleware', 'library'
    software_subtype VARCHAR(255),
    vendor VARCHAR(255),
    manufacturer VARCHAR(255),
    install_type VARCHAR(255),
    product_type VARCHAR(255),
    context TEXT,
    lifecycle_status VARCHAR(100),
    implementation_status VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create index for faster lookups
CREATE INDEX idx_software_display_name ON software(display_name);
CREATE INDEX idx_software_software_type ON software(software_type);
CREATE INDEX idx_software_vendor ON software(vendor);
CREATE INDEX idx_software_lifecycle_status ON software(lifecycle_status);

-- Create updated_at trigger for software table
CREATE TRIGGER update_software_timestamp 
BEFORE UPDATE ON software 
FOR EACH ROW 
EXECUTE FUNCTION update_timestamp();

-- Add a comment to document the table
COMMENT ON TABLE software IS 'Stores software application information'; 