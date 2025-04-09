-- migrations/1.sql
-- Application Portfolio Management (APM) Database Initialization

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create ENUM types
CREATE TYPE user_role AS ENUM ('organization_admin', 'application_portfolio_manager', 'stakeholder');
CREATE TYPE application_type AS ENUM ('application_software', 'system_software');
CREATE TYPE application_status AS ENUM ('active', 'deprecated', 'planned', 'under_development', 'retired');
CREATE TYPE entity_type AS ENUM ('organization', 'company', 'vendor', 'supplier');
CREATE TYPE log_action AS ENUM ('create', 'update', 'delete');

-- Organizations table
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    subdomain VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    avatar_url TEXT,
    mfa_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    mfa_secret VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Access logs
CREATE TABLE access_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    ip_address VARCHAR(45) NOT NULL,
    user_agent TEXT,
    role user_role NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Master categories
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Software types
CREATE TABLE software_types (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    parent_type_id UUID REFERENCES software_types(id) ON DELETE CASCADE,
    type application_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_name_per_parent UNIQUE (name, parent_type_id)
);

-- Master entities (vendors, companies, etc.)
CREATE TABLE master_entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    entity_type entity_type NOT NULL,
    description TEXT,
    website_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Organization-specific entities
CREATE TABLE organization_entities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    master_entity_id UUID NOT NULL REFERENCES master_entities(id) ON DELETE CASCADE,
    custom_name VARCHAR(255),
    custom_description TEXT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_master_entity_per_org UNIQUE (organization_id, master_entity_id)
);

-- Master applications
CREATE TABLE master_applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    vendor_id UUID REFERENCES master_entities(id) ON DELETE SET NULL,
    software_type_id UUID REFERENCES software_types(id) ON DELETE SET NULL,
    website_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_name_per_vendor UNIQUE (name, vendor_id)
);

-- Application-category relationships
CREATE TABLE master_application_categories (
    application_id UUID NOT NULL REFERENCES master_applications(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (application_id, category_id)
);

-- Competitor relationships
CREATE TABLE application_competitors (
    application_id UUID NOT NULL REFERENCES master_applications(id) ON DELETE CASCADE,
    competitor_id UUID NOT NULL REFERENCES master_applications(id) ON DELETE CASCADE,
    PRIMARY KEY (application_id, competitor_id),
    CHECK (application_id != competitor_id)
);

-- Organization applications
CREATE TABLE organization_applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    master_application_id UUID REFERENCES master_applications(id) ON DELETE SET NULL,
    custom_name VARCHAR(255) NOT NULL,
    custom_description TEXT,
    status application_status NOT NULL DEFAULT 'active',
    version VARCHAR(100),
    deployment_date DATE,
    end_of_life_date DATE,
    annual_cost DECIMAL(15, 2),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Application clusters
CREATE TABLE application_clusters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_cluster_name_per_org UNIQUE (organization_id, name)
);

-- Cluster-application relationships
CREATE TABLE cluster_applications (
    cluster_id UUID NOT NULL REFERENCES application_clusters(id) ON DELETE CASCADE,
    application_id UUID NOT NULL REFERENCES organization_applications(id) ON DELETE CASCADE,
    PRIMARY KEY (cluster_id, application_id)
);

-- News articles
CREATE TABLE news_articles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(500) NOT NULL,
    url TEXT NOT NULL UNIQUE,
    source VARCHAR(255) NOT NULL,
    published_date DATE NOT NULL,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- News article relationships
CREATE TABLE application_news (
    article_id UUID NOT NULL REFERENCES news_articles(id) ON DELETE CASCADE,
    application_id UUID NOT NULL REFERENCES master_applications(id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, application_id)
);

CREATE TABLE entity_news (
    article_id UUID NOT NULL REFERENCES news_articles(id) ON DELETE CASCADE,
    entity_id UUID NOT NULL REFERENCES master_entities(id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, entity_id)
);

-- Ranking sources
CREATE TABLE ranking_sources (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    url TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Application rankings
CREATE TABLE application_rankings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    application_id UUID NOT NULL REFERENCES master_applications(id) ON DELETE CASCADE,
    source_id UUID NOT NULL REFERENCES ranking_sources(id) ON DELETE CASCADE,
    rank INTEGER NOT NULL,
    ranking_date DATE NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_ranking_per_date UNIQUE (application_id, source_id, ranking_date)
);

-- Edit logs
CREATE TABLE edit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    entity_type VARCHAR(50) NOT NULL, -- 'organization', 'user', 'application', etc.
    entity_id UUID NOT NULL,
    action log_action NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    old_value TEXT,
    new_value TEXT,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_users_organization ON users(organization_id);
CREATE INDEX idx_access_logs_user ON access_logs(user_id);
CREATE INDEX idx_org_entities_organization ON organization_entities(organization_id);
CREATE INDEX idx_org_entities_master ON organization_entities(master_entity_id);
CREATE INDEX idx_master_applications_vendor ON master_applications(vendor_id);
CREATE INDEX idx_master_applications_type ON master_applications(software_type_id);
CREATE INDEX idx_org_applications_organization ON organization_applications(organization_id);
CREATE INDEX idx_org_applications_master ON organization_applications(master_application_id);
CREATE INDEX idx_clusters_organization ON application_clusters(organization_id);
CREATE INDEX idx_cluster_applications_cluster ON cluster_applications(cluster_id);
CREATE INDEX idx_application_news_application ON application_news(application_id);
CREATE INDEX idx_entity_news_entity ON entity_news(entity_id);
CREATE INDEX idx_rankings_application ON application_rankings(application_id);
CREATE INDEX idx_rankings_source ON application_rankings(source_id);
CREATE INDEX idx_edit_logs_user ON edit_logs(user_id);
CREATE INDEX idx_edit_logs_entity ON edit_logs(entity_type, entity_id);

-- Load initial categories from docs/categories.md
-- This could be handled by a separate migration script or application code

-- Create timestamp trigger function for updated_at
CREATE OR REPLACE FUNCTION update_timestamp() 
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for updated_at
CREATE TRIGGER update_organizations_timestamp BEFORE UPDATE ON organizations FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_users_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_categories_timestamp BEFORE UPDATE ON categories FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_software_types_timestamp BEFORE UPDATE ON software_types FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_master_entities_timestamp BEFORE UPDATE ON master_entities FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_organization_entities_timestamp BEFORE UPDATE ON organization_entities FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_master_applications_timestamp BEFORE UPDATE ON master_applications FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_organization_applications_timestamp BEFORE UPDATE ON organization_applications FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_application_clusters_timestamp BEFORE UPDATE ON application_clusters FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER update_ranking_sources_timestamp BEFORE UPDATE ON ranking_sources FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- Create function for search functionality
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Add comments to the schema
COMMENT ON TABLE users IS 'Users who can log in to the system';
COMMENT ON TABLE organizations IS 'Organizations with their own application portfolios';
COMMENT ON TABLE access_logs IS 'Log of user accesses to the system';
COMMENT ON TABLE categories IS 'Master list of software categories/functionalities';
COMMENT ON TABLE software_types IS 'Hierarchical types and subtypes of software';
COMMENT ON TABLE master_entities IS 'Master record of organizations, vendors, and other entities';
COMMENT ON TABLE organization_entities IS 'Organization-specific view of entities';
COMMENT ON TABLE master_applications IS 'Master record of all known applications';
COMMENT ON TABLE master_application_categories IS 'Categories that apply to each master application';
COMMENT ON TABLE application_competitors IS 'Competitive relationships between applications';
COMMENT ON TABLE organization_applications IS 'Applications within an organization''s portfolio';
COMMENT ON TABLE application_clusters IS 'Groups of related applications within an organization';
COMMENT ON TABLE cluster_applications IS 'Applications that belong to each cluster';
COMMENT ON TABLE news_articles IS 'News articles related to applications and entities';
COMMENT ON TABLE application_news IS 'Relationships between news articles and applications';
COMMENT ON TABLE entity_news IS 'Relationships between news articles and entities';
COMMENT ON TABLE ranking_sources IS 'Sources of application rankings';
COMMENT ON TABLE application_rankings IS 'Rankings of applications from various sources';
COMMENT ON TABLE edit_logs IS 'Audit log of all changes to the system'; 