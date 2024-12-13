# Groupie-Tracker
{{if eq .OnePieceBonus ""}}
                <h3 class="artifact-bonus">Two Pieces Bonus: {{.TwoPiecesBonus}}</h3>
                <h3 class="artifact-bonus">Four Pieces Bonus: {{.FourPiecesBonus}}</h3>
            {{else}}
                <h3 class="artifact-bonus">One Piece Bonus: {{.OnePieceBonus}}</h3>
            {{end}}