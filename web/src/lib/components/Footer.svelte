<script>
    import { SiteDescription, SiteTitle } from '$lib/info';
    import rssLogo from '$lib/assets/rss_logo.png';
</script>

<footer>
    <span>
        <p>{SiteTitle}</p>
        <p class="chill">
            <a href="mailto:ajones53.aj@gmail.com">ajones53.aj@gmail.com</a>
        </p>
    </span>

    <span>
        <p>
            <a href="https://github.com/ajone239">My Github</a>
        </p>
        <p class="chill">{SiteDescription}</p>
    </span>
</footer>

<style lang="scss">
    footer {
        padding: 1rem;
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        border-top: 2px dashed lightgrey;

        @media (max-width: 700px) {
            flex-direction: column;
        }
    }
    p.chill {
        opacity: 0.5;
    }
    a {
        color: inherit;
    }
</style>
