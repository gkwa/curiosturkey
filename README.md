# curiosturkey

Purpose: A tool to list Git repositories sorted by their most recent commit date.

## Example usage

```bash
# List repositories newer than 2 years
curiosturkey newerthan ~/daggerverse_repos /tmp/new_repos --age=2y

# List repositories newer than 2 days, hiding the age information
curiosturkey newerthan ~/daggerverse_repos --age=2d --hide-age

# List repositories newer than 1 year, 2 months, and 3 days
curiosturkey newerthan ~/projects ~/work/repos /opt/external_repos --age=1y2M3d

# List repositories newer than 6 hours
curiosturkey newerthan ~/quick_projects --age=6h

# List repositories newer than 3 weeks and 2 days, with verbose output
curiosturkey newerthan ~/long_term_projects --age=3w2d --verbose
