<script lang="ts">
    import MobileHeader from '$lib/components/MobileHeader.svelte';
    import DesktopHeader from '$lib/components/DesktopHeader.svelte';
    import Footer from '$lib/components/Footer.svelte';

    let { children } = $props();

    // Fake links
    const menuLinks = [
        { title: 'Blog', url: '/app' },
        { title: 'Arcade', url: '/app' },
        { title: 'About', url: '/app' }
    ];
</script>

<div class="site-wrapper">
    <div class="header-display-desktop">
        <DesktopHeader {menuLinks} />
    </div>

    <div class="header-display-mobile">
        <MobileHeader {menuLinks} />
    </div>

    <div class="main-content-wrapper">
        {@render children()}
    </div>

    <Footer />
</div>

<style>
    .site-wrapper {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
        flex-direction: column;
    }

    .main-content-wrapper {
        flex: 1;
        max-width: 45rem;
        width: 90%;
        margin: 0 auto;
        padding: 1rem;
    }

    .header-display-desktop {
        display: flex;
    }

    .header-display-mobile {
        display: none;
    }

    @media only screen and (pointer: coarse) and (max-width: 1024px),
        screen and (max-width: 799px) {
        .header-display-desktop {
            display: none;
        }

        .header-display-mobile {
            display: flex;
        }
    }
</style>
