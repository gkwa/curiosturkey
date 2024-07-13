# curiosturkey

Purpose: A tool to list Git repositories sorted by their most recent commit date.

## Example usage

```bash
# List repositories with commits made more recent than 2 years ago
curiosturkey newerthan ~/daggerverse_repos /tmp/new_repos --age=2y

# List repositories with commits made more recent than 2 days ago, hiding the age information
curiosturkey newerthan ~/daggerverse_repos --age=2d --hide-age

# List repositories with commits made more recent than 1 year, 2 months, and 3 days ago
curiosturkey newerthan ~/projects ~/work/repos /opt/external_repos --age=1y2M3d

# List repositories with commits made more recent than 6 hours ago
curiosturkey newerthan ~/quick_projects --age=6h

# List repositories with commits made more recent than 3 weeks and 2 days ago, with verbose output
curiosturkey newerthan ~/long_term_projects --age=3w2d --verbose
