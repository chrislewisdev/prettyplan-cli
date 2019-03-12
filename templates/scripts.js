function accordion(element) {
    const changes = element.parentElement.getElementsByClassName('changes');
    for (var i = 0; i < changes.length; i++) {
        changes[i].classList.toggle('collapsed');
    }
}

function expandAll() {
    const sections = document.querySelectorAll('.changes.collapsed');

    for (var i = 0; i < sections.length; i++) {
        sections[i].classList.toggle('collapsed');
    }

    document.querySelector('.expand-all').classList.toggle('hidden');
    document.querySelector('.collapse-all').classList.toggle('hidden');
}

function collapseAll() {
    const sections = document.querySelectorAll('.changes:not(.collapsed)');

    for (var i = 0; i < sections.length; i++) {
        sections[i].classList.toggle('collapsed');
    }

    document.querySelector('.expand-all').classList.toggle('hidden');
    document.querySelector('.collapse-all').classList.toggle('hidden');
}

function toggleRawPlan() {
    document.getElementById('raw-output').classList.toggle('collapsed');
}