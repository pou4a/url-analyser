export interface SubmitURLProps {
    onClose: () => void;
}

export interface URLType {
    id: number;
    url: string;
    status: string;
    created_at: string;
    updated_at: string;
}

export interface CrawlResult {
    id: number;
    url_id: number;
    html_version: string;
    title: string;
    h1_count: number;
    h2_count: number;
    h3_count: number;
    internal_links: number;
    external_links: number;
    inaccessible_links: number;
    has_login_form: boolean;
    analyzed_at: string;
}
